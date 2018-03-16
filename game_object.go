package spritengine

// GameObject represents a sprite and its properties
type GameObject struct {
	Sprite       SpriteInterface
	Position     Vector
	Mass         float64
	Velocity     Vector
	JumpVelocity float64
	Direction    int
	Flipped      bool
}

// IsJumping determined whether the game object is currently jumping
func (gameObject *GameObject) IsJumping(floorYPosition int) bool {

	// Special case for floating objects
	if gameObject.Mass == 0 {
		return false
	}

	return (int(gameObject.Position.Y) + gameObject.Height()) < floorYPosition

}

// Width gets the width of the game object
func (gameObject *GameObject) Width() int {

	return gameObject.Sprite.Width()

}

// Height gets the height of the game object
func (gameObject *GameObject) Height() int {

	return gameObject.Sprite.Height()

}

// RecalculatePosition recalculates the latest X and Y position of the game object from its properties
func (gameObject *GameObject) RecalculatePosition(gravity float64, floorYPosition int) {

	// Move left or right
	if gameObject.Direction == DirRight {
		gameObject.Position.X += gameObject.Velocity.X
	} else if gameObject.Direction == DirLeft {
		gameObject.Position.X -= gameObject.Velocity.X
	}

	// Jump up (and/or be pulled down by gravity)
	gameObject.Position.Y -= gameObject.Velocity.Y
	gameObject.Velocity.Y -= (gravity * gameObject.Mass)

	// Ensure the floor object acts as a barrier
	if (gameObject.Position.Y + float64(gameObject.Height())) > float64(floorYPosition) {
		gameObject.Position.Y = float64(floorYPosition - gameObject.Height())
	}

}
