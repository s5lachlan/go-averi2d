package main

import (
	"embed"
	_ "embed"
	"errors"
	"go-averi2d/sysimp"
	"strings"
)

//go:embed all:resource
var Assets embed.FS
var TextureAtlas sysimp.Texture2D
var SkidSound sysimp.Sound
var JumpSound sysimp.Sound
var ChooseSound sysimp.Sound
var SelectSound sysimp.Sound
var RunLeftSound sysimp.Sound
var RunRightSound sysimp.Sound

func LoadTextureFromMemory(path string) sysimp.Texture2D {
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
	image := sysimp.LoadImageFromMemory("."+filetype[len(filetype)-1], bytes, int32(len(bytes)))
	defer sysimp.UnloadImage(image)
	// Return the texture
	return sysimp.LoadTextureFromImage(image)
}

func LoadSoundFromMemory(path string) sysimp.Sound {
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
	wave := sysimp.LoadWaveFromMemory("."+filetype[len(filetype)-1], bytes, int32(len(bytes)))
	defer sysimp.UnloadWave(wave)
	return sysimp.LoadSoundFromWave(wave)
}
