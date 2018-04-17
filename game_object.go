package spritengine

// GameObject represents a sprite and its properties
type GameObject struct {
	CurrentState     string
	States           GameObjectStates
	Position         Vector
	Mass             float64
	Velocity         Vector
	Direction        int
	IsFlipped        bool
	IsControllable   bool
	IsFloor          bool
	IsInteractive    bool
	IsHidden         bool
	Level            *Level
	DynamicData      DynamicData
	FloorY           float64
	EventHandler     EventHandler
	CollisionHandler CollisionHandler
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
				return spriteSeries.Sprites[spriteCounter]
			}

		}

	}

	return spriteSeries.Sprites[0]

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

		wasAboveFloor := gameObject.Position.Y > gameObject.FloorY

		gameObject.Position.Y += gameObject.Velocity.Y
		gameObject.Velocity.Y -= (gravity * gameObject.Mass)

		// If actively falling down, emit the 'freefall' event
		if gameObject.Position.Y > gameObject.FloorY && gameObject.Velocity.Y < 0 {
			gameObject.EventHandler(EventFreefall, gameObject)
		}

		// Ensure the floor object acts as a barrier
		if gameObject.Position.Y <= gameObject.FloorY && gameObject.Mass != 0 {

			gameObject.Position.Y = gameObject.FloorY
			gameObject.Velocity.Y = 0

			if wasAboveFloor == true {
				gameObject.EventHandler(EventFloorCollision, gameObject)
			}

		}

	}

	// Don't fall too far off-screen
	if gameObject.Mass != 0 {

		minYPos := float64(0 - gameObject.Height())

		if gameObject.Position.Y <= minYPos {

			gameObject.Position.Y = minYPos

			if gameObject.IsInteractive == true {
				gameObject.EventHandler(EventDropOffLevel, gameObject)
			}

		}

	}

}

// SetDynamicData sets a piece of dynamic game object data
func (gameObject *GameObject) SetDynamicData(key string, value interface{}) {

	gameObject.DynamicData[key] = value

}

// GetDynamicData gets a piece of dynamic game object data, falling back to a
// defined value if the data does not exist
func (gameObject *GameObject) GetDynamicData(key string, fallback interface{}) interface{} {

	if value, ok := gameObject.DynamicData[key]; ok {
		return value
	}

	return fallback

}

// HasDynamicData checks whether a piece of dynamic game object data exists
func (gameObject *GameObject) HasDynamicData(key string) bool {

	if _, ok := gameObject.DynamicData[key]; ok {
		return true
	}

	return false

}

// ClearDynamicData clears a piece of dynamic game object data
func (gameObject *GameObject) ClearDynamicData(key string) {

	delete(gameObject.DynamicData, key)

}

// GetCollisionEdge infers the edge on which an intersecting object collided
// with the game object
func (gameObject *GameObject) GetCollisionEdge(collidingObject *GameObject) string {

	// Where is the game object's outer edge in relation to the colliding
	// object?
	isLeft := gameObject.Position.X < collidingObject.Position.X
	isRight := (gameObject.Position.X + float64(gameObject.Width())) > (collidingObject.Position.X + float64(collidingObject.Width()))
	isBottom := gameObject.Position.Y < collidingObject.Position.Y
	isTop := (gameObject.Position.Y + float64(gameObject.Height())) > (collidingObject.Position.Y + float64(collidingObject.Height()))

	// If both objects are at rest a simple 'left' or 'right' can be assumed,
	// regardless of the height of either object
	if (gameObject.Mass == 0 || gameObject.IsResting() == true) &&
		(collidingObject.Mass == 0 || collidingObject.IsResting() == true) {

		if isLeft == true {
			return EdgeLeft
		}

		if isRight == true {
			return EdgeRight
		}

	}

	// If there's any vertical movement, combination edges will make more sense
	if isLeft == true && isTop == true {
		return EdgeTopLeft
	}

	if isRight == true && isTop == true {
		return EdgeTopRight
	}

	if isLeft == true && isBottom == true {
		return EdgeBottomLeft
	}

	if isRight == true && isBottom == true {
		return EdgeBottomRight
	}

	// Failing the above, it may be the case that the game object is smaller
	// than the colliding object, and therefore fits nicely within one of the
	// regular edges
	if isLeft {
		return EdgeLeft
	}

	if isRight {
		return EdgeRight
	}

	if isBottom {
		return EdgeBottom
	}

	if isTop {
		return EdgeTop
	}

	// Should never reach here unless the game object is wholly within the
	// colliding object
	return EdgeNone

}
