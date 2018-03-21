package spritengine

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"strings"
)

// Sprite is a struct that represents sprite objects
type Sprite struct {
	Palette   *Palette
	Scanlines *[]int
}

// CreateSprite creates a sprite object based on a set of hex-encoded scanlines
func CreateSprite(palette *Palette, scanlines []int) (*Sprite, error) {

	if len(scanlines) != 32 {
		return nil, errors.New("Sprite not represented by the 32 hex groups required")
	}

	return &Sprite{
		Palette:   palette,
		Scanlines: &scanlines,
	}, nil

}

// AddToCanvas draws the sprite on an existing image canvas
func (sprite *Sprite) AddToCanvas(canvas *image.RGBA, targetX int, targetY int, mirrorImage bool) {

	// Return early if the sprite coordinates are off-canvas
	if targetX+sprite.Width() < 0 || targetX > canvas.Bounds().Max.X ||
		targetY+sprite.Height() < 0 || targetY > canvas.Bounds().Max.Y {
		return
	}

	spriteImage := image.NewRGBA(image.Rect(0, 0, 16, 16))

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

			xPos := xOffset + x

			if mirrorImage == true {
				xPos = (15 - xPos)
			}

			spriteImage.Set(xPos, y, (*sprite.Palette)[scanlinePixel])

		}

	}

	draw.Draw(canvas, spriteImage.Bounds().Add(image.Pt(targetX, targetY)), spriteImage, image.ZP, draw.Over)

}

// Width gets the pixel width of the sprite
func (sprite *Sprite) Width() int {

	return 16

}

// Height gets the pixel height of the sprite
func (sprite *Sprite) Height() int {

	return 16

}
