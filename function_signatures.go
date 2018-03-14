package spritengine

import (
	"image"

	"golang.org/x/mobile/event/key"
)

// FramePainter is the signature for functions that handle frame painting
type FramePainter func(stage *image.RGBA, frameRate float64)

// KeyListener is the signature for functions that handle key events
type KeyListener func(key.Event)
