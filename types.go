package spritengine

import (
	"image"
	"image/color"

	"golang.org/x/mobile/event/key"
)

// FramePainter is the signature for functions that handle frame painting
type FramePainter func(stage *image.RGBA, level *Level, frameRate float64)

// KeyListener is the signature for functions that handle key events for
// controllable game objects
type KeyListener func(event key.Event, gameObject *GameObject)

// EventHandler is the signature for functions that handle game events
type EventHandler func(eventCode int, gameObject *GameObject)

// CollisionHandler is the signature for functions that handle collision events
type CollisionHandler func(gameObject *GameObject, collision Collision)

// BeforePaint is the signature for functions that are called on levels prior
// to them being repainted
type BeforePaint func(level *Level)

// Vector is a struct to represent X/Y vectors
type Vector struct {
	X float64
	Y float64
}

// SpriteSeries is a type that defines a series of sprites that form an
// animation for a game object state
type SpriteSeries struct {
	Sprites         []SpriteInterface
	CyclesPerSecond int
}

// SpriteInterface is an interface that defines objects that can be treated a
// single sprites
type SpriteInterface interface {
	AddToCanvas(canvas *image.RGBA, targetX int, targetY int, mirrorImage bool)
	Width() int
	Height() int
}

// Palette is a type that defines sprite palettes
type Palette map[string]color.RGBA

// GameObjectStates is a type that defines the various states (and accompanying
// sprites) of game objects
type GameObjectStates map[string]SpriteSeries

// DynamicData is a type that defines a repository of arbitrary game object
// data
type DynamicData map[string]interface{}

// Collision is a struct that represents a collision with another game object
type Collision struct {
	GameObject *GameObject
	Edge       string
}
