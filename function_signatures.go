package spritengine

import (
	"image"

	"golang.org/x/mobile/event/key"
)

// FramePainter is the signature for functions that handle frame painting
type FramePainter func(stage *image.RGBA, level *Level, frameRate float64)

// KeyListener is the signature for functions that handle key events for
// controllable game objects
type KeyListener func(event key.Event, gameObject *GameObject)
