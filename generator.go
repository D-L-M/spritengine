package spritengine

import (
	"image"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// GenerateFromPNGFile generates a SpriteGroup package file from an image on disk
func GenerateFromPNGFile(inputFile string, outputFile string, packageName string, exportedSpriteName string) {

	imgFile, imgFileErr := os.Open(inputFile)
	imgFileConfig, imgFileConfErr := os.Open(inputFile)

	if imgFileErr != nil || imgFileConfErr != nil {
		log.Fatal("Error reading input file")
	}

	defer imgFile.Close()
	defer imgFileConfig.Close()

	img, _, imgErr := image.Decode(imgFile)
	config, _, confErr := image.DecodeConfig(imgFileConfig)

	if imgErr != nil || confErr != nil {
		log.Fatal("Error reading image file")
	}

	if config.Width%16 != 0 || config.Height%16 != 0 {
		log.Fatal("The image width and/or height was not a multiple of 16")
	}

	// Figure out the palette first
	tempPalette := map[string]color.RGBA{}

	for x := 0; x < config.Width; x++ {

		for y := 0; y < config.Height; y++ {

			colourString, colourRGBA := getColourStringAndRGBA(img.At(x, y))
			tempPalette[colourString] = colourRGBA

		}

	}

	if len(tempPalette) > 16 {
		log.Fatal("More than 16 colours used in image palette")
	}

	palette := map[string]color.RGBA{}
	paletteSlots := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	slotCount := 0

	for _, colour := range tempPalette {
		palette[paletteSlots[slotCount]] = colour
		slotCount++
	}

	// Get the palette slots for each pixel
	sprites := map[string][]string{}

	for y := 0; y < config.Height; y++ {

		spriteCountY := int(math.Floor(float64(y) / 16))

		for x := 0; x < config.Width; x++ {

			spriteCountX := int(math.Floor(float64(x) / 16))
			spriteCoord := strconv.Itoa(spriteCountX) + "," + strconv.Itoa(spriteCountY)
			_, colourRGBA := getColourStringAndRGBA(img.At(x, y))

			for paletteSlot, paletteColour := range palette {

				if paletteColour == colourRGBA {
					sprites[spriteCoord] = append(sprites[spriteCoord], paletteSlot)
				}

			}

		}

	}

	// Build up the output file
	spriteStrings := []string{}
	spriteNames := []string{}

	// Palette
	paletteName := "palette_" + packageName
	paletteString := "var " + paletteName + " = &map[string]color.RGBA{"

	for paletteSlot, colour := range palette {
		paletteString += `"` + paletteSlot + `": color.RGBA{` + strconv.Itoa(int(colour.R)) + `,` + strconv.Itoa(int(colour.G)) + `,` + strconv.Itoa(int(colour.B)) + `,` + strconv.Itoa(int(colour.A)) + `},`
	}

	paletteString += "}"

	// Sprites
	for y := 0; y < (config.Height / 16); y++ {

		for x := 0; x < (config.Width / 16); x++ {

			spriteName := "sprite_" + packageName + "_" + strconv.Itoa(x) + "_" + strconv.Itoa(y)
			spriteString := "var " + spriteName + ", _ = spritengine.CreateSprite(" + paletteName + ", []int{"
			spriteNames = append(spriteNames, spriteName)

			for i, paletteSlot := range sprites[strconv.Itoa(x)+","+strconv.Itoa(y)] {

				if (i % 8) == 0 {
					spriteString += "0x"
				}

				spriteString += paletteSlot

				if (i % 8) == 7 {
					spriteString += ","
				}

			}

			spriteString += "})"

			spriteStrings = append(spriteStrings, spriteString)

		}

	}

	fileContents := `package ` + packageName + `

import (
	"image/color"

	"github.com/D-L-M/spritengine"
)

` + paletteString + `

` + strings.Join(spriteStrings, "\n") + `

var ` + exportedSpriteName + `, _ = spritengine.CreateSpriteGroup(` + strconv.Itoa(config.Width/16) + `, ` + strconv.Itoa(config.Height/16) + `, &[]*spritengine.Sprite{` + strings.Join(spriteNames, ",") + `})`

	// Write the output file to disk
	err := ioutil.WriteFile(outputFile, []byte(fileContents), 0644)

	if err != nil {
		log.Fatal("Error writing to ourput file")
	}

}

// getColourStringAndRGBA converts a color.Color object to its string and color.RGBA representations
func getColourStringAndRGBA(colour color.Color) (string, color.RGBA) {

	r, g, b, a := colour.RGBA()
	colourString := strconv.Itoa(int(r)) + "," + strconv.Itoa(int(g)) + "," + strconv.Itoa(int(b)) + "," + strconv.Itoa(int(a))
	colourRGBA := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

	return colourString, colourRGBA

}
