package main

import "go-averi2d/sysimp"

func Flip(boolean bool) float32 {
	if boolean {
		return -1
	} else {
		return 1
	}
}

func UnloadSounds(sounds ...sysimp.Sound) {
	for _, sound := range sounds {
		sysimp.UnloadSound(sound)
	}
}

func ApproachZero(arg, amt int) int {
	if arg < -amt {
		return arg + amt
	}
	if arg > amt {
		return arg - amt
	}
	return 0
}

// Ternary operator: condition ? value1 : value2
func Ternary(condition bool, value1 any, value2 any) any {
	if condition {
		return value1
	} else {
		return value2
	}
}

func ScaleDisplay() {
	var screenWidthFloat float32 = float32(sysimp.GetScreenWidth())
	var screenHeightFloat float32 = float32(sysimp.GetScreenHeight())

	// ADJUST DISPLAY AREA in case window is resized
	if screenWidthFloat/screenHeightFloat > MaxAspectRatio {
		// If the screen is now wider than max_aspect_ratio: scale by height, shift display right to be centered in the window
		ScaleFactor = screenHeightFloat / IdealScreenHeight
		DisplayX = int32(
			(screenWidthFloat - (2.5 * screenHeightFloat)) / 2,
		)
		DisplayY = 0
	} else if screenWidthFloat/screenHeightFloat < MinAspectRatio {
		// If the screen is now taller than it is wide: scale by width, shift display left to cut off as much as min_aspect_ratio permits
		ScaleFactor = screenWidthFloat / (IdealScreenHeight * MinAspectRatio)
		DisplayX = int32(
			(screenWidthFloat - (IdealScreenWidth * ScaleFactor)) / 2,
		)
		DisplayY = int32(
			(screenHeightFloat - screenWidthFloat) / 2,
		) // ...and shift display down to be centered
	} else {
		// If the screen is between the min and max aspect ratios, scale by height and shift left to cut off the sides as needed
		ScaleFactor = screenHeightFloat / IdealScreenHeight
		DisplayX = int32(
			(screenWidthFloat - (IdealScreenWidth * ScaleFactor)) / 2,
		)
		DisplayY = 0
	}

}
