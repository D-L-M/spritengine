package spritengine

import (
	"image"
	"image/color"
	"image/draw"
)

// World is a struct that defines a single world of a game
type World struct {
	BackgroundColour color.RGBA
	Gravity          float64
	GameObjects      []*GameObject
	Game             *Game
}

// Repaint redraws the entire world for a new frame
func (world *World) Repaint(stage *image.RGBA) {

	// Add background
	draw.Draw(stage, stage.Bounds(), &image.Uniform{color.RGBA{0, 138, 197, 255}}, image.ZP, draw.Src)

	// Update each game object
	for _, gameObject := range world.GameObjects {

		gameObject.RecalculatePosition(world.Gravity, world.Game.Height)

		if gameObject.Direction == DirLeft {
			gameObject.Flipped = true
		} else if gameObject.Direction == DirRight {
			gameObject.Flipped = false
		}

		gameObject.Sprite.AddToCanvas(stage, int(gameObject.Position.X), int(gameObject.Position.Y), gameObject.Flipped)

	}

}
