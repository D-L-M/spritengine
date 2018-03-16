package spritengine

import (
	"image"
	"image/color"
	"image/draw"
)

// Level is a struct that defines a single level of a game
type Level struct {
	BackgroundColour color.RGBA
	Gravity          float64
	GameObjects      []*GameObject
	Game             *Game
}

// Repaint redraws the entire level for a new frame
func (level *Level) Repaint(stage *image.RGBA) {

	// Add background
	draw.Draw(stage, stage.Bounds(), &image.Uniform{color.RGBA{0, 138, 197, 255}}, image.ZP, draw.Src)

	// Update each game object
	for _, gameObject := range level.GameObjects {

		gameObject.RecalculatePosition(level.Gravity, level.Game.Height)

		if gameObject.Direction == DirLeft {
			gameObject.Flipped = true
		} else if gameObject.Direction == DirRight {
			gameObject.Flipped = false
		}

		gameObject.Sprite.AddToCanvas(stage, int(gameObject.Position.X), int(gameObject.Position.Y), gameObject.Flipped)

	}

}
