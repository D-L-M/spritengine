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

	// Paint the background colour
	draw.Draw(stage, stage.Bounds(), &image.Uniform{level.BackgroundColour}, image.ZP, draw.Src)

	// Update each game object
	for _, gameObject := range level.GameObjects {

		gameObject.Level = level
		gameObject.RecalculatePosition(level.Gravity)

		if gameObject.Direction == DirLeft {
			gameObject.Flipped = true
		} else if gameObject.Direction == DirRight {
			gameObject.Flipped = false
		}

		// 0 is at the bottom, so flip the Y axis to paint correctly
		invertedY := level.Game.Height - int(gameObject.Position.Y) - gameObject.Height()

		gameObject.CurrentSprite().AddToCanvas(stage, int(gameObject.Position.X), invertedY, gameObject.Flipped)

	}

}
