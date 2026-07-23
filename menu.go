package main

import rl "github.com/gen2brain/raylib-go/raylib"

var MenuHighlightTime int
var MenuSelection MenuSelectionEnum

type MenuSelectionEnum int

const MenuOffsetY = 200

const (
	MenuStartGame MenuSelectionEnum = iota
	MenuOptions
	MenuQuit
)

type Menu struct {
	BaseScene // Embed default scene stuff
}

// Creates a new menu button with the text included as a child
func NewMenuButton(menuSelection MenuSelectionEnum) Scene {
	offset := float32(menuSelection)
	return &BaseScene{
		SourceX: 82, SourceY: 2564,
		SourceWidth: 112, SourceHeight: 17, Scale: menuScale,
		DestX: 550, DestY: MenuOffsetY + ((20 * offset) * menuScale), Visible: true, Render: true,
		Children: []Scene{
			&BaseScene{
				SourceX: 82, SourceY: 2564 + (17 * (offset + 1)),
				SourceWidth: 55, SourceHeight: 17, Scale: 1.0,
				DestX: 550, DestY: MenuOffsetY + ((20 * offset) * menuScale), Visible: true, Render: true,
			},
		},
	}
}

func NewMenu() *Menu {
	baseScene := BaseScene{
		Render: false,
		DestX:  40, DestY: 40, Scale: 1.0,
		Visible: true,
		Paused:  false,
		Children: []Scene{
			NewMenuButton(MenuStartGame),
			NewMenuButton(MenuOptions),
			NewMenuButton(MenuQuit),
			MenuHighlightButton,
		},
	}
	return &Menu{baseScene}
}

func (m *Menu) Draw() {
	m.BaseScene.Draw()
	rl.DrawText("Averi 2D Platformer", int32(530*ScaleFactor+float32(DisplayX)), int32(50*ScaleFactor+float32(DisplayY)), int32(74*ScaleFactor), rl.DarkGreen)
	rl.DrawText("Recreational Go Port", int32(530*ScaleFactor+float32(DisplayX)), int32(120*ScaleFactor+float32(DisplayY)), int32(34*ScaleFactor), rl.DarkGreen)
}

func (m *Menu) Update() {
	// This counts up every 20 frames and goes back down to 0
	if MenuHighlightTime >= 20 {
		MenuHighlightTime = 0
	} else {
		MenuHighlightTime += 1
	}

	// This controls the menu, W to move up, S to go down, and limit between 0 and 2.
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyS) {
		rl.SetSoundVolume(ChooseSound, 0.5)
		rl.PlaySound(ChooseSound)
	}
	if rl.IsKeyPressed(rl.KeyW) && MenuSelection > 0 {
		MenuSelection -= 1
	} else if rl.IsKeyPressed(rl.KeyS) && MenuSelection < 2 {
		MenuSelection += 1
	}

	// This makes the highlight button flash for 10 frames when the highlight timer is above 10 frames.
	MenuHighlightButton.DestY = 200 + (20 * float32(MenuSelection) * menuScale)
	if MenuHighlightTime > 10 {
		MenuHighlightButton.SourceX = 82 + 14
	} else {
		MenuHighlightButton.SourceX = 82
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		rl.SetSoundVolume(SelectSound, 0.5)
		rl.PlaySound(SelectSound)
		switch MenuSelection {
		case MenuStartGame:
			rl.SetExitKey(rl.KeyNull)
			CurrentScene = GameScene
		case MenuQuit:
			ExitGame = true
		}
	}
}
