package spritengine

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"strings"
)

// Sprite is a struct that represents sprite objects
type Sprite struct {
	Palette   *map[string]color.RGBA
	Scanlines *[]int
}

// CreateSprite creates a sprite object based on a set of hex-encoded scanlines
func CreateSprite(palette *map[string]color.RGBA, scanlines []int) (*Sprite, error) {

	if len(scanlines) != 32 {
		return nil, errors.New("Sprite not represented by the 32 hex groups required")
	}

	return &Sprite{
		Palette:   palette,
		Scanlines: &scanlines,
	}, nil

}

// AddToCanvas draws the sprite on an existing image canvas
func (sprite *Sprite) AddToCanvas(canvas *image.RGBA, targetX int, targetY int) {

	for i, scanline := range *sprite.Scanlines {

		y := i
		xOffset := 0

		if (i % 2) != 0 {
			y--
			xOffset = 8
		}

		y /= 2

		scanlineString := fmt.Sprintf("%08x", scanline)
		scanlinePixels := strings.Split(scanlineString, "")

		for x, scanlinePixel := range scanlinePixels {
			canvas.Set((targetX + xOffset + x), (targetY + y), (*sprite.Palette)[scanlinePixel])
		}

	}

}
