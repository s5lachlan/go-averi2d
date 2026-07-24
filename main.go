// A Go rewrite of averi2d platformerby
// Rewrite by Lachlan, original by Flareonkek

package main

import "go-averi2d/sysimp"

var FrameCount int

func main() {
	sysimp.SetConfigFlags(sysimp.FlagWindowResizable)

	sysimp.InitWindow(int32(IdealScreenWidth), int32(IdealScreenHeight), "goAveri2D")
	sysimp.InitAudioDevice()
	defer sysimp.CloseAudioDevice()
	defer sysimp.CloseWindow()
	// Load spritesheet stored in the executable
	TextureAtlas = LoadTextureFromMemory("resource/spritesheet.png")
	defer sysimp.UnloadTexture(TextureAtlas)
	SkidSound = LoadSoundFromMemory("resource/skidding.wav")
	JumpSound = LoadSoundFromMemory("resource/jump.wav")
	ChooseSound = LoadSoundFromMemory("resource/choose.wav")
	SelectSound = LoadSoundFromMemory("resource/select.wav")
	RunLeftSound = LoadSoundFromMemory("resource/runleft.wav")
	RunRightSound = LoadSoundFromMemory("resource/runright.wav")
	defer UnloadSounds(SkidSound, JumpSound, ChooseSound, SelectSound, RunLeftSound, RunRightSound)

	sysimp.SetTargetFPS(30)

	for !ExitGame && !sysimp.WindowShouldClose() {
		sysimp.BeginDrawing()
		sysimp.ClearBackground(BackgroundColour)

		ScaleDisplay()
		// Hardcoded background and floor for now
		sysimp.DrawRectangle(DisplayX, DisplayY, int32(ScaleFactor*IdealScreenWidth), int32(ScaleFactor*IdealScreenHeight), sysimp.GetColor(0x90EE90ff))
		sysimp.DrawRectangleV(
			sysimp.NewVector2(0+float32(DisplayX), 350*ScaleFactor+float32(DisplayY)),
			sysimp.NewVector2(IdealScreenWidth*ScaleFactor, (IdealScreenHeight-350)*ScaleFactor),
			sysimp.Gray,
		)

		CurrentScene.Update()
		CurrentScene.Draw()
		sysimp.EndDrawing()
		FrameCount += 1
	}

}
