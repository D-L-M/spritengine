package spritengine

import (
	"golang.org/x/mobile/event/key"
)

// GameObject represents a sprite and its properties
type GameObject struct {
	States          GameObjectStates
	Position        Vector
	Mass            float64
	CurrentVelocity Vector
	SlowVelocity    Vector
	FastVelocity    Vector
	Direction       int
	Flipped         bool
	Controllable    bool
	FastMoving      bool
	Level           *Level
}

// IsJumping determined whether the game object is currently jumping
func (gameObject *GameObject) IsJumping() bool {

	// TODO: Change this so that it acts on objects beneath the game object
	floorYPosition := 0

	// Special case for floating objects
	if gameObject.Mass == 0 {
		return false
	}

	return int(gameObject.Position.Y) > floorYPosition

}

// CurrentSprite gets the current sprite for the object's state
func (gameObject *GameObject) CurrentSprite() SpriteInterface {

	spriteSeries := gameObject.States["default"]

	if gameObject.Direction == DirLeft || gameObject.Direction == DirRight {
		spriteSeries = gameObject.States["moving"]
	}

	if gameObject.IsJumping() {
		spriteSeries = gameObject.States["jumping"]
	}

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
	gameObject.Position.Y += gameObject.CurrentVelocity.Y
	gameObject.CurrentVelocity.Y -= (gravity * gameObject.Mass)

	// Ensure the floor object acts as a barrier
	if gameObject.Position.Y < float64(floorYPosition) {
		gameObject.Position.Y = float64(floorYPosition)
	}

}

// ActOnInputEvent repositions the game object based on an input event
func (gameObject *GameObject) ActOnInputEvent(event key.Event) {

	switch event.Code {

	case key.CodeLeftArrow:

		if event.Direction == key.DirPress && gameObject.IsJumping() == false && gameObject.Direction == DirStationary {
			gameObject.Direction = DirLeft
		} else if event.Direction == key.DirRelease && gameObject.Direction == DirLeft {
			gameObject.Direction = DirStationary
		}

	case key.CodeRightArrow:

		if event.Direction == key.DirPress && gameObject.IsJumping() == false && gameObject.Direction == DirStationary {
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

		if event.Direction == key.DirPress && gameObject.IsJumping() == false {

			if gameObject.FastMoving == true {
				gameObject.CurrentVelocity.Y = gameObject.FastVelocity.Y
			} else {
				gameObject.CurrentVelocity.Y = gameObject.SlowVelocity.Y
			}

		}

	}

}
