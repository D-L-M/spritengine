package spritengine

import (
	"golang.org/x/mobile/event/key"
)

// GameObject represents a sprite and its properties
type GameObject struct {
	Sprite          SpriteInterface
	Position        Vector
	Mass            float64
	CurrentVelocity Vector
	SlowVelocity    Vector
	FastVelocity    Vector
	Direction       int
	Flipped         bool
	Controllable    bool
	FastMoving      bool
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

	xVelocity := gameObject.SlowVelocity.X

	if gameObject.FastMoving == true {
		xVelocity = gameObject.FastVelocity.X
	}

	// Move left or right
	if gameObject.Direction == DirRight {
		gameObject.Position.X += xVelocity
	} else if gameObject.Direction == DirLeft {
		gameObject.Position.X -= xVelocity
	}

	// Jump up (and/or be pulled down by gravity)
	gameObject.Position.Y -= gameObject.CurrentVelocity.Y
	gameObject.CurrentVelocity.Y -= (gravity * gameObject.Mass)

	// Ensure the floor object acts as a barrier
	if (gameObject.Position.Y + float64(gameObject.Height())) > float64(floorYPosition) {
		gameObject.Position.Y = float64(floorYPosition - gameObject.Height())
	}

}

// ActOnInputEvent repositions the game object based on an input event
func (gameObject *GameObject) ActOnInputEvent(event key.Event) {

	// TODO: Change this so that it acts on objects beneath the game object
	height := 240

	switch event.Code {

	case key.CodeLeftArrow:

		if event.Direction == key.DirPress && gameObject.IsJumping(height) == false && gameObject.Direction == DirStationary {
			gameObject.Direction = DirLeft
		} else if event.Direction == key.DirRelease && gameObject.Direction == DirLeft {
			gameObject.Direction = DirStationary
		}

	case key.CodeRightArrow:

		if event.Direction == key.DirPress && gameObject.IsJumping(height) == false && gameObject.Direction == DirStationary {
			gameObject.Direction = DirRight
		} else if event.Direction == key.DirRelease && gameObject.Direction == DirRight {
			gameObject.Direction = DirStationary
		}

	case key.CodeLeftShift, key.CodeRightShift:

		if event.Direction == key.DirPress || event.Direction == key.Direction(10) {
			gameObject.FastMoving = true
		} else if event.Direction == key.DirRelease || event.Direction == key.Direction(11) {
			gameObject.FastMoving = false
		}

	case key.CodeSpacebar:

		if event.Direction == key.DirPress && gameObject.IsJumping(height) == false {

			if gameObject.FastMoving == true {
				gameObject.CurrentVelocity.Y = gameObject.FastVelocity.Y
			} else {
				gameObject.CurrentVelocity.Y = gameObject.SlowVelocity.Y
			}

		}

	}

}
