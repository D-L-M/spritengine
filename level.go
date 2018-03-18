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

	// Figure out where all the floor objects are
	level.AssignFloors()

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

// AssignFloors iterates through all objects in the level and defines which
// object beneath them (if any) should be considered their 'floor' object,
// setting its top edge as the lowest point that the object can fall
func (level *Level) AssignFloors() {

	floorXCoords := map[int][]*GameObject{}

	// Make a map of each object's possible X positions
	for _, gameObject := range level.GameObjects {

		for i := 0; i < gameObject.Width(); i++ {

			xPos := i + int(gameObject.Position.X)
			floorXCoords[xPos] = append(floorXCoords[xPos], gameObject)

		}

	}

	// Find the objects that sit beneath every other object
	for _, gameObject := range level.GameObjects {

		// Skip objects that float
		if gameObject.Mass == 0 {
			continue
		}

		highestFloorObject := 0.0

		for i := 0; i < gameObject.Width(); i++ {

			xPos := i + int(gameObject.Position.X)

			if floorObjects, ok := floorXCoords[xPos]; ok {

				// Find the one that is highest while still being lower than
				// the object itself
				for _, floorObject := range floorObjects {

					floorObjectTop := (floorObject.Position.Y + float64(floorObject.Height()))

					if floorObjectTop <= gameObject.Position.Y {

						if floorObjectTop > highestFloorObject {
							highestFloorObject = floorObjectTop
						}

					}

				}

			}

		}

		gameObject.FloorY = highestFloorObject

	}

}
