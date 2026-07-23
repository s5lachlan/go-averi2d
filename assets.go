package main

import (
	"embed"
	_ "embed"
	"errors"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed all:resource
var Assets embed.FS
var TextureAtlas rl.Texture2D
var SkidSound rl.Sound
var JumpSound rl.Sound
var ChooseSound rl.Sound
var SelectSound rl.Sound
var RunLeftSound rl.Sound
var RunRightSound rl.Sound

func LoadTextureFromMemory(path string) rl.Texture2D {
	// Check for a file extension
	filetype := strings.Split(path, ".")
	if len(filetype) <= 1 {
		panic(errors.New("Filetype is required to load from memory."))
	}
	// Read bytes from file
	bytes, err := Assets.ReadFile(path)
	if err != nil {
		panic(err)
	}
	// Load image from the bytes, using the length of the bytes array as it's size and
	// the type of the file from filetype variable from earlier
	image := rl.LoadImageFromMemory("."+filetype[len(filetype)-1], bytes, int32(len(bytes)))
	defer rl.UnloadImage(image)
	// Return the texture
	return rl.LoadTextureFromImage(image)
}

func LoadSoundFromMemory(path string) rl.Sound {
	// Check for a file extension
	filetype := strings.Split(path, ".")
	if len(filetype) <= 1 {
		panic(errors.New("Filetype is required to load from memory."))
	}
	// Read bytes from file
	bytes, err := Assets.ReadFile(path)
	if err != nil {
		panic(err)
	}
	wave := rl.LoadWaveFromMemory("."+filetype[len(filetype)-1], bytes, int32(len(bytes)))
	defer rl.UnloadWave(wave)
	return rl.LoadSoundFromWave(wave)
}
