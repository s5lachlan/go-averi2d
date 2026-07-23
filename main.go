// A Go rewrite of averi2d platformerby
// Rewrite by Lachlan, original by Flareonkek

package main

import rl "github.com/gen2brain/raylib-go/raylib"

var FrameCount int

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(int32(IdealScreenWidth), int32(IdealScreenHeight), "goAveri2D")
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
	// Load spritesheet stored in the executable
	TextureAtlas = LoadTextureFromMemory("resource/spritesheet.png")
	defer rl.UnloadTexture(TextureAtlas)
	SkidSound = LoadSoundFromMemory("resource/skidding.wav")
	JumpSound = LoadSoundFromMemory("resource/jump.wav")
	ChooseSound = LoadSoundFromMemory("resource/choose.wav")
	SelectSound = LoadSoundFromMemory("resource/select.wav")
	RunLeftSound = LoadSoundFromMemory("resource/runleft.wav")
	RunRightSound = LoadSoundFromMemory("resource/runright.wav")
	defer UnloadSounds(SkidSound, JumpSound, ChooseSound, SelectSound, RunLeftSound, RunRightSound)

	rl.SetTargetFPS(30)

	for !ExitGame && !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(BackgroundColour)

		ScaleDisplay()
		// Hardcoded background and floor for now
		rl.DrawRectangle(DisplayX, DisplayY, int32(ScaleFactor*IdealScreenWidth), int32(ScaleFactor*IdealScreenHeight), rl.GetColor(0x90EE90ff))
		rl.DrawRectangleV(
			rl.NewVector2(0+float32(DisplayX), 350*ScaleFactor+float32(DisplayY)),
			rl.NewVector2(IdealScreenWidth*ScaleFactor, (IdealScreenHeight-350)*ScaleFactor),
			rl.Gray,
		)

		CurrentScene.Update()
		CurrentScene.Draw()
		rl.EndDrawing()
		FrameCount += 1
	}

}
