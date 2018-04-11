package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/D-L-M/spritengine"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/event/key"
)

// Entrypoint to the example game
func main() {

	levels := []*spritengine.Level{
		getLevel(),
	}

	spritengine.CreateGame("Game Example", 320, 224, 2, 30, framePainter, keyListener, levels)

}

// framePainter adds additional graphics to the painted level frame
func framePainter(stage *image.RGBA, level *spritengine.Level, frameRate float64) {

	writeText(stage, "FPS: "+fmt.Sprintf("%.2f", frameRate), 10, 20)

}

// writeText writes text to the stage
func writeText(stage *image.RGBA, text string, xPos int, yPos int) {

	fontDrawer := font.Drawer{
		Dst:  stage,
		Src:  image.NewUniform(color.RGBA{0, 0, 0, 255}),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(xPos * 64), Y: fixed.Int26_6(yPos * 64)},
	}

	fontDrawer.DrawString(text)

}

// keyListener reacts to key events for controllable game objects
func keyListener(event key.Event, gameObject *spritengine.GameObject) {

	switch event.Code {

	case key.CodeLeftArrow:

		if event.Direction == key.DirPress && gameObject.IsResting() == true && gameObject.Direction == spritengine.DirStationary {
			gameObject.Direction = spritengine.DirLeft
			gameObject.CurrentState = "moving"
		} else if event.Direction == key.DirRelease && gameObject.Direction == spritengine.DirLeft {
			gameObject.Direction = spritengine.DirStationary
			gameObject.CurrentState = "standing"
		}

	case key.CodeRightArrow:

		if event.Direction == key.DirPress && gameObject.IsResting() == true && gameObject.Direction == spritengine.DirStationary {
			gameObject.Direction = spritengine.DirRight
			gameObject.CurrentState = "moving"
		} else if event.Direction == key.DirRelease && gameObject.Direction == spritengine.DirRight {
			gameObject.Direction = spritengine.DirStationary
			gameObject.CurrentState = "standing"
		}

	case key.CodeSpacebar:

		if event.Direction == key.DirPress && gameObject.IsResting() == true {
			gameObject.CurrentState = "jumping"
			gameObject.Velocity.Y = 6
		}

	}

}

// getLevel gets the example level
func getLevel() *spritengine.Level {

	return &spritengine.Level{
		Gravity:          0.5,
		BackgroundColour: color.RGBA{200, 200, 200, 255},
		BeforePaint:      beforePaint,
		GameObjects:      []*spritengine.GameObject{},
		PaintOffset: spritengine.Vector{
			X: 0,
			Y: 0,
		},
	}

}

// getMaxScrollX gets the maximum level scroll width
func getMaxScrollX() float64 {

	return 1600.0

}

// beforePaint handles minor reworkings of the level prior to repainting
func beforePaint(level *spritengine.Level) {

	// Make the camera follow the controllable game object
	for _, gameObject := range level.GameObjects {

		if gameObject.IsControllable {

			xOffset := gameObject.Position.X - float64(level.Game.Width/2)

			if xOffset < 0 {
				xOffset = 0
			}

			if xOffset > getMaxScrollX() {
				xOffset = getMaxScrollX()
			}

			level.PaintOffset.X = xOffset

		}

	}

}
