package spritengine

import "image"

// SpriteInterface is an interface that defines objects that can be treated a single sprites
type SpriteInterface interface {
	AddToCanvas(canvas *image.RGBA, targetX int, targetY int, mirrorImage bool)
	Width() int
	Height() int
}
