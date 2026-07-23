package main

import (
	"fmt"
	"math/rand"

	_ "github.com/gen2brain/raylib-go/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	BaseScene            // Embed default scene stuff
	Speed                float32
	RunSpeedMax          float32
	VelocityX, VelocityY float32
	TravelSinceLastWalk  float32
	TailSwingDirection   int32
}

const (
	AveriIdle            = -1
	AveriWalkStart       = 0
	AveriRunStart        = 2
	AveriRunEnd          = 9
	AveriJumpRising      = 15
	AveriJumpPeak        = 16
	AveriJumpDescent     = 17
	AveriJumpFastDescent = 18
)

const AverageStride = 16

func NewPlayer() *Player {
	baseScene := BaseScene{
		Render:  true,
		SourceX: 0, SourceY: 90,
		FrameY: -1, FrameX: 0,
		FrameYAmount: 90, FrameXAmount: 0,
		SourceWidth: 81, SourceHeight: 90,
		DestX: 40, DestY: 40, Scale: 1.0,
		Visible: true,
		Paused:  false,
		Children: []Scene{
			&BaseScene{
				Render:  true,
				SourceX: 180, SourceY: 35,
				SourceWidth: 45, SourceHeight: 35,
				DestX: 500, DestY: 40,
				Scale:          1.0,
				Visible:        false,
				FlipHorizontal: true,
			},
		},
	}
	return &Player{
		BaseScene:          baseScene,
		RunSpeedMax:        15.0,
		Speed:              500.0,
		TailSwingDirection: 1,
	}
}

func (self *Player) Draw() {
	Tail := self.Children[0].Get()
	self.BaseScene.Draw()
	rl.DrawText(
		fmt.Sprintf(
			`Debug mode
X: %f, Y: %f
VelocityX: %f, VelocityY: %f
FrameY: %d
FlipHorizontal: %b
TravelSinceLastWalk: %f`,
			self.DestX, self.DestY,
			self.VelocityX, self.VelocityY,
			self.FrameY,
			self.FlipHorizontal,
			self.TravelSinceLastWalk), 0, 0, int32(32*ScaleFactor), rl.DarkGreen)
	// Jumping animations
	if self.DestY < float32(HardcodedGroundValue) {
		if self.VelocityY < -2 {
			self.FrameY = AveriJumpRising
		} else if self.VelocityY < 2 {
			self.FrameY = AveriJumpPeak
		} else if self.VelocityY < 6 {
			self.FrameY = AveriJumpFastDescent
		} else {
			self.FrameY = AveriJumpFastDescent
		}
	} else if self.FlipHorizontal && self.VelocityX < 0 || !self.FlipHorizontal && self.VelocityX > 0 {
		// Sliding and turning around
		if self.FrameY < 10 || self.FrameY > 14 {
			rl.SetSoundPitch(SkidSound, 0.9+rl.Clamp(rand.Float32(), 0.0, 0.1))
			rl.PlaySound(SkidSound)
			self.FrameY = 10
		} else if self.FrameY != 14 && FrameCount%3 == 0 {
			self.FrameY += 1
		}
		// When you're about to have slid to a halt...
		if self.VelocityX == 1 || self.VelocityX == -1 {
			self.VelocityX = -self.VelocityX
			self.FrameY = AveriRunStart
		}
	} else {
		// General walking/running/standing
		if self.TravelSinceLastWalk < -AverageStride || self.TravelSinceLastWalk > AverageStride {
			// Every time you moved AverageStride pixels, advance the animation frame and update TravelSinceLastWalk
			self.TravelSinceLastWalk = float32(ApproachZero(int(self.TravelSinceLastWalk), AverageStride))
			self.FrameY += 1
			if FrameCount%9 == 0 {
				rl.PlaySound(RunLeftSound)
			}
			if FrameCount%5 == 0 {
				rl.PlaySound(RunRightSound)
			}
		}
		// If we reach the end of the run cycle, go back to its start
		if self.FrameY > AveriRunEnd {
			self.FrameY = AveriRunStart
		}
		// If she isn't moving, go back to stand-still sprite and reset TravelSinceLastWalk
		if self.VelocityX == 0 {
			self.FrameY = AveriIdle
			self.TravelSinceLastWalk = 0
		} else if self.FrameY == AveriIdle { // But if she is moving even a bit,
			Tail.Visible = false
			self.FrameY = AveriWalkStart // don't wait for animDX to kickstart the walk animation
		}
	}
	// Swing tail if on ground and not moving
	if self.VelocityX == 0 && self.DestY == float32(HardcodedGroundValue) {
		// Swing Tail
		Tail.Visible = true
		Tail.FlipHorizontal = self.FlipHorizontal
		if Tail.FlipHorizontal {
			Tail.DestX = self.DestX
		} else {
			Tail.DestX = self.DestX + 36
		}
		Tail.DestY = self.DestY + 44

		if FrameCount%7 == 0 {
			if Tail.FrameY+1 == 6 || Tail.FrameY+1 == 0 {
				self.TailSwingDirection = -self.TailSwingDirection
			}
			Tail.FrameY += 1 * self.TailSwingDirection
		}

		if Tail.FrameY > 5 {
			Tail.FrameY = -1
		} else if Tail.FrameY < -1 {
			Tail.FrameY = 5
		}
	} else {
		Tail.Visible = false
	}
}

func (self *Player) Update() {
	if rl.IsKeyPressed(rl.KeyX) {
		if self.Children[0].Get().FrameY > 4 {
			self.Children[0].Get().FrameY = 0
		}
		self.Children[0].Get().FrameY += 1
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		rl.SetExitKey(rl.KeyEscape)
		CurrentScene = MenuScene
	}
	if rl.IsKeyDown(rl.KeyA) {
		if self.VelocityX > -self.RunSpeedMax && (FrameCount%3 != 0) {
			self.VelocityX -= 1
		}
		self.FlipHorizontal = false
	} else if rl.IsKeyDown(rl.KeyD) {
		if self.VelocityX < self.RunSpeedMax && (FrameCount%3 != 0) {
			self.VelocityX += 1
		}
		self.FlipHorizontal = true

	} else if FrameCount%3 != 0 {
		self.VelocityX = float32(ApproachZero(int(self.VelocityX), 1))
	}

	self.DestX += self.VelocityX
	self.TravelSinceLastWalk += self.VelocityX

	if self.DestY < float32(HardcodedGroundValue) {
		if rl.IsKeyDown(rl.KeySpace) {
			self.VelocityY += 1
		} else {
			self.VelocityY += 3
		}
	} else if rl.IsKeyDown(rl.KeySpace) {
		rl.SetSoundVolume(JumpSound, 0.3)
		rl.SetSoundPitch(JumpSound, 0.9+rl.Clamp(rand.Float32(), 0.0, 0.1))
		rl.PlaySound(JumpSound)
		self.VelocityY = -20
	}

	if self.DestY+self.VelocityY >= float32(HardcodedGroundValue) {
		self.DestY = float32(HardcodedGroundValue)
		self.VelocityY = 0
	} else {
		self.DestY += self.VelocityY
	}
}
