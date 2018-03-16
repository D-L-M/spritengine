package spritengine

// SpriteSeries is a type that defines a series of sprites that form an
// animation for a game object state
type SpriteSeries struct {
	Sprites         []SpriteInterface
	CyclesPerSecond int
}
