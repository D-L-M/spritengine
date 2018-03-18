package spritengine

// GameObject represents a sprite and its properties
type GameObject struct {
	CurrentState string
	States       GameObjectStates
	Position     Vector
	Mass         float64
	Velocity     Vector
	Direction    int
	Flipped      bool
	Controllable bool
	Level        *Level
	DynamicData  DynamicData
	FloorY       float64
	Interactive  bool
}

// IsResting determined whether the game object is currently atop another game
// object
func (gameObject *GameObject) IsResting() bool {

	// Special case for floating objects
	if gameObject.Mass == 0 {
		return false
	}

	return int(gameObject.Position.Y) == int(gameObject.FloorY)

}

// CurrentSprite gets the current sprite for the object's state
func (gameObject *GameObject) CurrentSprite() SpriteInterface {

	spriteSeries := gameObject.States[gameObject.CurrentState]
	sprite := gameObject.getCurrentSpriteFrame(spriteSeries)

	return sprite

}

// getCurrentSpriteFrame gets the appropriate frame of a sprite series based on the
// game's frame ticker
func (gameObject *GameObject) getCurrentSpriteFrame(spriteSeries SpriteSeries) SpriteInterface {

	spriteIndex := 0

	if gameObject.Level != nil {

		game := gameObject.Level.Game
		framesPerSprite := (game.TargetFrameRate / spriteSeries.CyclesPerSecond) / len(spriteSeries.Sprites)
		spriteCounter := 0
		i := 0

		for j := 0; j < game.TargetFrameRate; j++ {

			i++

			if i == framesPerSprite {
				i = 0
				spriteCounter++
			}

			if spriteCounter >= len(spriteSeries.Sprites) {
				spriteCounter = 0
			}

			if j == game.CurrentFrame {
				spriteIndex = spriteCounter
			}

		}

	}

	return spriteSeries.Sprites[spriteIndex]

}

// Width gets the width of the game object
func (gameObject *GameObject) Width() int {

	return gameObject.CurrentSprite().Width()

}

// Height gets the height of the game object
func (gameObject *GameObject) Height() int {

	return gameObject.CurrentSprite().Height()

}

// RecalculatePosition recalculates the latest X and Y position of the game
// object from its properties
func (gameObject *GameObject) RecalculatePosition(gravity float64) {

	// Move left or right
	if gameObject.Direction == DirRight {
		gameObject.Position.X += gameObject.Velocity.X
	} else if gameObject.Direction == DirLeft {
		gameObject.Position.X -= gameObject.Velocity.X
	}

	// Jump up (and/or be pulled down by gravity) if the floor is further down
	if gameObject.FloorY <= gameObject.Position.Y {
		gameObject.Position.Y += gameObject.Velocity.Y
		gameObject.Velocity.Y -= (gravity * gameObject.Mass)
	}

	// Ensure the floor object acts as a barrier
	if gameObject.Position.Y < gameObject.FloorY {

		gameObject.Position.Y = gameObject.FloorY
		gameObject.Velocity.Y = 0

		// TODO: Make a method that can be called that the user provides on
		// events like this so they can choose to update the state
		if gameObject.Controllable == true {

			if gameObject.Direction == DirStationary {
				gameObject.CurrentState = "standing"
			} else {
				gameObject.CurrentState = "moving"
			}

		}

	}

	// Only fall just off-screen if not floating
	if gameObject.Mass != 0 {

		minYPos := (0 - float64(gameObject.Height()))

		if gameObject.Position.Y <= minYPos {

			gameObject.Position.Y = minYPos

			// Mark as non-interactive
			// TODO: Move this out to a collision event with the absolute floor (0 - height)
			gameObject.Interactive = false

		}

	}

}
