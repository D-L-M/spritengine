package spritengine

// GameObject represents a sprite and its properties
type GameObject struct {
	Sprite       SpriteInterface
	Position     Vector
	Width        int
	Height       int
	Mass         int
	Velocity     Vector
	JumpVelocity int
	Direction    int
}

// IsJumping determined whether the game object is currently jumping
func (gameObject *GameObject) IsJumping(floorYPosition int) bool {

	return (gameObject.Position.Y + gameObject.Height) < floorYPosition

}

// RecalculatePosition recalculates the latest X and Y position of the game object from its properties
func (gameObject *GameObject) RecalculatePosition(gravity int, floorYPosition int) {

	// Move left or right
	if gameObject.Direction == DirRight {
		gameObject.Position.X += gameObject.Velocity.X
	} else if gameObject.Direction == DirLeft {
		gameObject.Position.X -= gameObject.Velocity.X
	}

	// Jump up (and/or be pulled down by gravity)
	gameObject.Position.Y -= gameObject.Velocity.Y
	gameObject.Velocity.Y -= gravity

	// Ensure the floor object acts as a barrier
	if (gameObject.Position.Y + gameObject.Height) > floorYPosition {
		gameObject.Position.Y = (floorYPosition - gameObject.Height)
	}

}
