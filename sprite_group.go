package spritengine

import (
	"errors"
	"image"
	"strconv"
)

// SpriteGroup is a struct that represents a group of sprites that form a larger individual sprite
type SpriteGroup struct {
	Width   int
	Height  int
	Sprites *[]*Sprite
}

// AddToCanvas draws the sprite group on an existing image canvas
func (spriteGroup *SpriteGroup) AddToCanvas(canvas *image.RGBA, targetX int, targetY int, mirrorImage bool) {

	spriteCount := 0

	for y := 0; y < spriteGroup.Height; y++ {

		if mirrorImage == true {

			for x := (spriteGroup.Width - 1); x >= 0; x-- {

				(*spriteGroup.Sprites)[spriteCount].AddToCanvas(canvas, (targetX + (x * 16)), (targetY + (y * 16)), mirrorImage)

				spriteCount++

			}

		} else {

			for x := 0; x < spriteGroup.Width; x++ {

				(*spriteGroup.Sprites)[spriteCount].AddToCanvas(canvas, (targetX + (x * 16)), (targetY + (y * 16)), mirrorImage)

				spriteCount++

			}

		}

	}

}

// CreateSpriteGroup creates a sprite group based on a grid size and collection of sprites
func CreateSpriteGroup(width int, height int, sprites *[]*Sprite) (*SpriteGroup, error) {

	if len(*sprites) != (width * height) {
		return nil, errors.New("Sprite group requires " + strconv.Itoa(width*height) + " sprites, not " + strconv.Itoa(len(*sprites)))
	}

	return &SpriteGroup{
		Width:   width,
		Height:  height,
		Sprites: sprites,
	}, nil

}
