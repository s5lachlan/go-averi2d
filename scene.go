package main

import "go-averi2d/sysimp"

type Scene interface {
	Update()
	Draw()
	Get() *BaseScene
}

type BaseScene struct {
	// Describe a rectangle on the sprite sheet, what part of it to draw
	SourceX, SourceY,
	SourceWidth, SourceHeight float32

	// When these values are set, the Source X and Source Y are automatically
	// adjusted, so SourceY = 90 and FrameY = 1, moves down to the next Y frame.
	FrameX, FrameY, FrameXAmount, FrameYAmount int32

	// Describe a location on the screen on which to draw that sprite
	DestX, DestY float32

	Scale          float32 // A scene's own scale factor
	Visible        bool    // Defines if the scene is visible
	Render         bool    // Defines if the scene should render
	Paused         bool    // Defines if the scene is paused
	FlipHorizontal bool    // Defines if the scene if flipped horizontally
	FlipVertical   bool    // Definesif the scene is flipped vertically
	Children       []Scene
}

func (s *BaseScene) Update() {
	if !s.Paused {
		for _, child := range s.Children {
			if !child.Get().Paused {
				child.Update()
			}
		}
	}
}
func (s *BaseScene) Draw() {
	if s.Visible {
		for _, child := range s.Children {
			childScene := child.Get()
			if childScene.Visible {
				if childScene.Render {
					sysimp.DrawTexturePro(
						TextureAtlas,
						sysimp.NewRectangle(
							childScene.SourceX*float32(1+childScene.FrameX),
							childScene.SourceY*float32(1+childScene.FrameY),
							childScene.SourceWidth*Flip(childScene.FlipHorizontal),
							childScene.SourceHeight,
						),

						sysimp.NewRectangle(
							childScene.DestX*ScaleFactor+float32(DisplayX),
							childScene.DestY*ScaleFactor+float32(DisplayY),
							childScene.SourceWidth*ScaleFactor*(childScene.Scale*s.Scale),
							childScene.SourceHeight*ScaleFactor*(childScene.Scale*s.Scale),
						),
						sysimp.Vector2Zero(),
						0.0,
						sysimp.White,
					)
				}
				child.Draw()
			}
		}
	}
}
func (s *BaseScene) Get() *BaseScene {
	return s
}

func CreateScene(children ...Scene) Scene {
	newScene := &BaseScene{Visible: true, Scale: 1.0}
	newScene.Children = append(newScene.Children, children...)
	return newScene
}
