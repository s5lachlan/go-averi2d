// This is the globals file, where all globally available variables are put so it's easier to keep track of them.

package main

import "go-averi2d/sysimp"

// Screen resolution and aspect ratio
const IdealScreenWidth float32 = 1800 // The sides will be cut off on less wide screens (basically all screens)
const IdealScreenHeight float32 = 720 // But all 720 height units will always be visible

const MaxAspectRatio float32 = 2.5 // This is also the "ideal" aspect ratio
const MinAspectRatio float32 = 1.0

var BackgroundColour = sysimp.GetColour(0x052c46ff)
var ScaleFactor float32 = 1.0

var DisplayX, DisplayY int32

// Scenes
var GameScene = CreateScene(
	NewPlayer(),
)
var MenuScene = NewMenu()

// Current Scene, decides what the current scene is. To switch scenes do:
// CurrentScene = NewScene
var CurrentScene Scene = MenuScene

// Main Menu variables

// Menu Scale
var menuScale float32 = 2.0

// Decides if the game should close
var ExitGame bool

// Menu Highlight
var MenuHighlightButton = &BaseScene{
	SourceX: 82, SourceY: 2564 - 17,
	SourceWidth: 16, SourceHeight: 17, Scale: menuScale,
	DestX: 550, DestY: MenuOffsetY * menuScale, Visible: true, Render: true,
}

// Game variables
var HardcodedGroundValue = 260
