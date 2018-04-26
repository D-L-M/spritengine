package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/D-L-M/spritengine"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/event/key"
)

// Entrypoint to the example game
func main() {

	levels := []*spritengine.Level{
		getLevel(),
	}

	spritengine.CreateGame("Game Example", 320, 224, 2, 64, framePainter, keyListener, levels)

}

// framePainter adds additional graphics to the painted level frame
func framePainter(stage *image.RGBA, level *spritengine.Level, frameRate float64) {

	// Framerate
	writeText(stage, "FPS: "+fmt.Sprintf("%.2f", frameRate), 10, 20)

	// Progress
	writeText(stage, "Progress: "+fmt.Sprintf("%.0f", ((level.PaintOffset.X/getMaxScrollX())*100))+"%", 10, 35)

}

// writeText writes text to the stage
func writeText(stage *image.RGBA, text string, xPos int, yPos int) {

	fontDrawer := font.Drawer{
		Dst:  stage,
		Src:  image.NewUniform(color.RGBA{0, 0, 0, 255}),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(xPos * 64), Y: fixed.Int26_6(yPos * 64)},
	}

	fontDrawer.DrawString(text)

}

// keyListener reacts to key events for controllable game objects
func keyListener(event key.Event, gameObject *spritengine.GameObject) {

	switch event.Code {

	case key.CodeLeftArrow:

		if event.Direction == key.DirPress && gameObject.IsResting() == true && gameObject.Direction == spritengine.DirStationary {
			gameObject.Direction = spritengine.DirLeft
			gameObject.CurrentState = "moving"
		} else if event.Direction == key.DirRelease && gameObject.Direction == spritengine.DirLeft {
			gameObject.Direction = spritengine.DirStationary
			gameObject.CurrentState = "standing"
		}

	case key.CodeRightArrow:

		if event.Direction == key.DirPress && gameObject.IsResting() == true && gameObject.Direction == spritengine.DirStationary {
			gameObject.Direction = spritengine.DirRight
			gameObject.CurrentState = "moving"
		} else if event.Direction == key.DirRelease && gameObject.Direction == spritengine.DirRight {
			gameObject.Direction = spritengine.DirStationary
			gameObject.CurrentState = "standing"
		}

	case key.CodeSpacebar:

		if event.Direction == key.DirPress && gameObject.IsResting() == true {
			gameObject.CurrentState = "jumping"
			gameObject.Velocity.Y = 6
		}

	}

}

// getLevel gets the example level
func getLevel() *spritengine.Level {

	gameObjects := []*spritengine.GameObject{}

	// Clouds
	for i := 0; i < 8; i++ {
		cloudX := float64((i * 150) + rand.Intn(100-10) + 10)
		cloudY := float64(rand.Intn(200-150) + 150)
		gameObjects = append(gameObjects, getCloud(cloudX, cloudY))
	}

	// Floor
	for i := 0; i < 80; i++ {

		if i > 24 && i < 29 {
			continue
		}

		yPos := 0.0

		if i > 48 && i < 55 {
			yPos = 40.0
		}

		gameObjects = append(gameObjects, getFloor(float64(i*16), yPos))

	}

	for i := 58; i < 62; i++ {
		gameObjects = append(gameObjects, getFloor(float64(i*16), 85.0))
	}

	// Powerups
	gameObjects = append(gameObjects, getPowerup(950, 170))

	// Character
	gameObjects = append(gameObjects, getCharacter(20, 16))

	return &spritengine.Level{
		Gravity:          1,
		BackgroundColour: color.RGBA{126, 192, 238, 255},
		BeforePaint:      beforePaint,
		GameObjects:      gameObjects,
		PaintOffset: spritengine.Vector{
			X: 0,
			Y: 0,
		},
	}

}

// getMaxScrollX gets the maximum level scroll width
func getMaxScrollX() float64 {

	return 960.0

}

// beforePaint handles minor reworkings of the level prior to repainting
func beforePaint(level *spritengine.Level) {

	// Make the camera follow the controllable game object
	for _, gameObject := range level.GameObjects {

		if gameObject.IsControllable {

			xOffset := gameObject.Position.X - float64(level.Game.Width/2)
			yOffset := gameObject.Position.Y - float64(level.Game.Height/2)

			if xOffset < 0 {
				xOffset = 0
			}

			if xOffset > getMaxScrollX() {
				xOffset = getMaxScrollX()
			}

			if gameObject.Position.Y < float64(level.Game.Height/2) {
				yOffset = 0
			}

			level.PaintOffset.X = xOffset
			level.PaintOffset.Y = yOffset

		}

	}

}

// Sprite information for the controllable character
var paletteCharacter = &spritengine.Palette{"4": color.RGBA{0, 0, 0, 255}, "6": color.RGBA{97, 56, 53, 255}, "9": color.RGBA{46, 26, 35, 255}, "0": color.RGBA{227, 156, 118, 255}, "2": color.RGBA{69, 41, 51, 255}, "c": color.RGBA{110, 192, 155, 255}, "3": color.RGBA{84, 53, 56, 255}, "5": color.RGBA{65, 128, 121, 255}, "b": color.RGBA{129, 86, 70, 255}, "a": color.RGBA{138, 88, 61, 255}, "8": color.RGBA{0, 0, 0, 0}, "e": color.RGBA{88, 89, 76, 255}, "1": color.RGBA{255, 209, 164, 255}, "d": color.RGBA{189, 130, 89, 255}, "7": color.RGBA{36, 81, 87, 255}}

var spriteCharacterMoving100, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving110, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88444488, 0x88888888})
var spriteCharacterMoving101, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x84888888, 0x88888888, 0x46484444, 0x88888888, 0x4644666a, 0x88888888, 0x422266aa, 0x88888888, 0x42266aaa, 0x88888888, 0x466aaaa6, 0x88888888, 0x4266aa00, 0x88888888, 0x466a6029, 0x88888888, 0x422660dd, 0x88888888, 0x4106d069, 0x88888888, 0x4102d011, 0x88888888, 0x842dd111, 0x88888888, 0x88400111, 0x88888888, 0x84440011, 0x88888884, 0x45c5557a, 0x8888884c, 0xcc5c5557})
var spriteCharacterMoving111, _ = spritengine.CreateSprite(paletteCharacter, []int{0x4466aa48, 0x88888888, 0xaaa6aaa4, 0x88888888, 0xaaaaaaa4, 0x88888888, 0xaaaa6aa4, 0x88888888, 0xaaaa6a48, 0x88888888, 0x6a66a488, 0x88888888, 0x00004888, 0x88888888, 0x69096488, 0x88888888, 0xdaad4888, 0x88888888, 0x0ad94844, 0x48888888, 0x1da144ed, 0x48888888, 0x11117eee, 0x48888888, 0x001193e4, 0x88888888, 0x11199948, 0x88888888, 0xad9b9488, 0x88888888, 0x79ebb488, 0x44488888})
var spriteCharacterMoving102, _ = spritengine.CreateSprite(paletteCharacter, []int{0x888884cc, 0xccc555cc, 0x888884cc, 0xccc55ccc, 0x88888455, 0x5cc55cc6, 0x888884cc, 0xcc4556ed, 0x88888840, 0x1112bbdd, 0x88888841, 0x11099bbb, 0x88888884, 0x11199555, 0x88888888, 0x44994339, 0x88888888, 0x88444bbb, 0x88888888, 0x8844433b, 0x88888888, 0x84294333, 0x88888888, 0x4229bbb3, 0x88888888, 0x4243bbbb, 0x88888888, 0x84843b34, 0x88888888, 0x88884448, 0x88888888, 0x88888888})
var spriteCharacterMoving112, _ = spritengine.CreateSprite(paletteCharacter, []int{0xc6ddb444, 0x00048888, 0xedddc47d, 0x00048888, 0xdddcc57d, 0xa0048888, 0xddccc575, 0xd0048888, 0xdcccc557, 0x54488888, 0xcccc5555, 0x48888888, 0x55554444, 0x88888888, 0x9d194888, 0x88888888, 0xbbdb4888, 0x88888888, 0xdddd4888, 0x88888888, 0xbdddd488, 0x88888888, 0x3dddd488, 0x88888888, 0x44bddb48, 0x88888888, 0x884ddd48, 0x88888888, 0x884b9948, 0x88888888, 0x88849334, 0x88888888})
var characterMoving1, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving100, spriteCharacterMoving110, spriteCharacterMoving101, spriteCharacterMoving111, spriteCharacterMoving102, spriteCharacterMoving112})

var spriteCharacterMoving200, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving210, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving201, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x84888888, 0x88888888, 0x46484444, 0x88888888, 0x4644666a, 0x88888888, 0x422266aa, 0x88888888, 0x42266aaa, 0x88888888, 0x466aaaa6, 0x88888888, 0x4266aa00, 0x88888888, 0x466a6092, 0x88888888, 0x422660dd, 0x88888888, 0x4106d092, 0x88888888, 0x4112d011, 0x88888888, 0x842dd111, 0x88888888, 0x88400111, 0x88888888, 0x88440011, 0x88888888, 0x84c5557a})
var spriteCharacterMoving211, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88444488, 0x88888888, 0x4466aa48, 0x88888888, 0xaaa6aaa4, 0x88888888, 0xaaaaaaa4, 0x88888888, 0xaaaa6aa4, 0x88888888, 0xaaaa6a48, 0x88888888, 0x6a66a488, 0x88888888, 0x00004888, 0x88888888, 0x92029488, 0x88888888, 0xdadd4888, 0x88888888, 0x0da24888, 0x88888888, 0x1ad14844, 0x48888888, 0x111144ed, 0x48888888, 0x00112eee, 0x48888888, 0x111297e4, 0x88888888, 0xda7b9948, 0x88888888})
var spriteCharacterMoving202, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x4c5c5557, 0x88888884, 0xccc555cc, 0x8888884c, 0xccc55ccc, 0x8888884c, 0xc5c55cc6, 0x8888884c, 0x5c01166d, 0x88888884, 0x5c1111bd, 0x88888888, 0x441110bb, 0x88888888, 0x84211555, 0x88888888, 0x84999399, 0x88888888, 0x84994bbb, 0x88888888, 0x884443bd, 0x88888888, 0x84229333, 0x88888888, 0x84229bbb, 0x88888888, 0x842443b3, 0x88888888, 0x84488444, 0x88888888, 0x88888888})
var spriteCharacterMoving212, _ = spritengine.CreateSprite(paletteCharacter, []int{0x796bb488, 0x88888888, 0xc6ddb444, 0x48888888, 0x6dddc400, 0x04888888, 0xdddccd00, 0x04888888, 0xddcccdad, 0x04888888, 0xdcccc5d0, 0x04888888, 0xcccc5554, 0x48888888, 0x55554448, 0x88888888, 0xd0094888, 0x88888888, 0xbbdb4888, 0x88888888, 0xdddd4888, 0x88888888, 0xbddd4888, 0x88888888, 0x3ddb4888, 0x88888888, 0x49948888, 0x88888888, 0x43348888, 0x88888888, 0x43334888, 0x88888888})
var characterMoving2, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving200, spriteCharacterMoving210, spriteCharacterMoving201, spriteCharacterMoving211, spriteCharacterMoving202, spriteCharacterMoving212})

var spriteCharacterMoving300, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving310, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving301, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888844, 0x88888888, 0x84444466, 0x88888888, 0x46222666, 0x88888888, 0x422666aa, 0x88888888, 0x466aaaaa, 0x88888888, 0x42266aa0, 0x88888888, 0x4666a602, 0x88888888, 0x4222660d, 0x88888888, 0x42106d02, 0x88888888, 0x84112d01, 0x88888888, 0x8422dd11, 0x88888888, 0x884d0011, 0x88888888, 0x8884d001, 0x88888888, 0x8847557a, 0x88888888, 0x84755c55, 0x88888888, 0x8475c5c5})
var spriteCharacterMoving311, _ = spritengine.CreateSprite(paletteCharacter, []int{0x44444448, 0x88888888, 0x6aaa6aa4, 0x88888888, 0xaaaaa6aa, 0x48888888, 0xaaaaaaaa, 0x48888888, 0x66a666a4, 0x88888888, 0x00006648, 0x88888888, 0x92909488, 0x88888888, 0xddda4888, 0x88888888, 0x90aa4444, 0x88888888, 0x11da4ec4, 0x88888888, 0x1111eee4, 0x88888888, 0x1001ee48, 0x88888888, 0x11199348, 0x88888888, 0xda7b9488, 0x88888888, 0x7796b488, 0x88888888, 0x55edb488, 0x88888888})
var spriteCharacterMoving302, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x47755ccc, 0x88888888, 0x4775cccc, 0x88888888, 0x4da75ccc, 0x88888888, 0x4dddbdcc, 0x88888888, 0x4da42bbb, 0x88888888, 0x84429555, 0x88888888, 0x84299399, 0x88888888, 0x84999bbb, 0x88888888, 0x84994bdd, 0x88888888, 0x88444ddd, 0x88888888, 0x8884bddd, 0x88888888, 0x8884ddd4, 0x88888888, 0x884bdb44, 0x88888888, 0x88499484, 0x88888888, 0x88433488, 0x88888888, 0x88433348})
var spriteCharacterMoving312, _ = spritengine.CreateSprite(paletteCharacter, []int{0x56d10148, 0x88888888, 0xccc11114, 0x88888888, 0xc5c11114, 0x88888888, 0xc5c01148, 0x88888888, 0xb5555488, 0x88888888, 0x55554888, 0x88888888, 0xd0094888, 0x88888888, 0xbdb3b488, 0x88888888, 0xdd3bb488, 0x88888888, 0xd3bbb488, 0x88888888, 0x993b4888, 0x88888888, 0x22948888, 0x88888888, 0x22488888, 0x88888888, 0x24888888, 0x88888888, 0x48888888, 0x88888888, 0x88888888, 0x88888888})
var characterMoving3, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving300, spriteCharacterMoving310, spriteCharacterMoving301, spriteCharacterMoving311, spriteCharacterMoving302, spriteCharacterMoving312})

var spriteCharacterMoving400, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x84888444})
var spriteCharacterMoving410, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x84484444, 0x88888888, 0x4aa4aaa4, 0x88888888})
var spriteCharacterMoving401, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x46444666, 0x88888888, 0x4222666a, 0x88888884, 0x22666aaa, 0x88888884, 0x66aaaaa6, 0x88888884, 0x2266aa00, 0x88888884, 0x666a6069, 0x88888884, 0x222660dd, 0x88888884, 0x2106d029, 0x88888888, 0x4112d011, 0x88888888, 0x422dd111, 0x88888888, 0x84d00111, 0x88888888, 0x884d0011, 0x88888888, 0x8477557a, 0x88888888, 0x47755557, 0x88888884, 0x77755c5c, 0x8888884d, 0xdd75c5cc})
var spriteCharacterMoving411, _ = spritengine.CreateSprite(paletteCharacter, []int{0xaaa6aaa4, 0x88888888, 0xaaaaaaa4, 0x88888888, 0xaaaa6a48, 0x88888888, 0x6a666488, 0x88888888, 0x00044888, 0x88888888, 0x29094888, 0x88888888, 0xdda48888, 0x88888888, 0x0ad48888, 0x88888888, 0x1aa44488, 0x88888888, 0x111ed488, 0x88888888, 0x0013e488, 0x88888888, 0x11994888, 0x88888888, 0xa9b94888, 0x88888888, 0x79eb4444, 0x88888888, 0x556b4101, 0x48888888, 0xccc5c111, 0x14888888})
var spriteCharacterMoving402, _ = spritengine.CreateSprite(paletteCharacter, []int{0x8888884d, 0xada75ccc, 0x8888884d, 0xdad755cc, 0x88888884, 0xad466bdd, 0x88888888, 0x4442bbbb, 0x88888888, 0x88495555, 0x88888888, 0x8429399d, 0x88888888, 0x8499bbbb, 0x88888888, 0x8499bddd, 0x88888884, 0x444ddddd, 0x88888843, 0x39ddddd4, 0x88888843, 0x39bddb48, 0x88888843, 0x44444488, 0x88888844, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving412, _ = spritengine.CreateSprite(paletteCharacter, []int{0xccc5c111, 0x14888888, 0xccc5cc11, 0x48888888, 0xcccc5444, 0x88888888, 0xb5544888, 0x88888888, 0x55488888, 0x88888888, 0x10948888, 0x88888888, 0xdb3b4888, 0x88888888, 0xd3bbb488, 0x88888888, 0x33bbbb48, 0x88888888, 0x4443bb48, 0x88888888, 0x88849948, 0x88888888, 0x88842248, 0x88888888, 0x88842224, 0x88888888, 0x88884444, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var characterMoving4, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving400, spriteCharacterMoving410, spriteCharacterMoving401, spriteCharacterMoving411, spriteCharacterMoving402, spriteCharacterMoving412})

var spriteCharacterMoving500, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x84888888})
var spriteCharacterMoving510, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88444488, 0x88888888, 0x4466aa48, 0x88888888})
var spriteCharacterMoving501, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x46484444, 0x88888888, 0x4644666a, 0x88888884, 0x222266aa, 0x88888884, 0x22266aaa, 0x88888884, 0x666aaaa6, 0x88888884, 0x2266aa00, 0x88888884, 0x666a6092, 0x88888884, 0x222660dd, 0x88888884, 0x2106d092, 0x88888888, 0x4112d011, 0x88888888, 0x422dd111, 0x88888888, 0x84d00111, 0x88888888, 0x884d0011, 0x88888888, 0x884757da, 0x88888888, 0x84755577, 0x88888888, 0x47755555})
var spriteCharacterMoving511, _ = spritengine.CreateSprite(paletteCharacter, []int{0xaaa6aaa4, 0x88888888, 0xaaaaaaa4, 0x88888888, 0xaaaa6aa4, 0x88888888, 0xaaaa6a48, 0x88888888, 0x6a66a488, 0x88888888, 0x00044888, 0x88888888, 0x92024888, 0x88888888, 0xdad48888, 0x88888888, 0x0aa44448, 0x88888888, 0x1ad4e548, 0x88888888, 0x111eee48, 0x88888888, 0x0013e488, 0x88888888, 0x11999488, 0x88888888, 0xd9b94888, 0x88888888, 0x956b4888, 0x88888888, 0x55eb4444, 0x48888888})
var spriteCharacterMoving502, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x4d755c55, 0x88888884, 0xdad7c5cc, 0x88888884, 0xada75ccc, 0x88888888, 0x4dd6bccc, 0x88888888, 0x8429bbcc, 0x88888888, 0x42995555, 0x88888888, 0x4999399d, 0x88888888, 0x4994bbbb, 0x88888888, 0x4444bddd, 0x88888884, 0x3944dddd, 0x88888843, 0x39bddddd, 0x88888843, 0x4bddddd4, 0x88888884, 0x84bddb48, 0x88888888, 0x88444488, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving512, _ = spritengine.CreateSprite(paletteCharacter, []int{0x56bd4101, 0x14888888, 0x5555c111, 0x14888888, 0xccc5c111, 0x14888888, 0xccc5cc11, 0x48888888, 0xcccccc44, 0x88888888, 0x55544488, 0x88888888, 0x00948888, 0x88888888, 0xdd348888, 0x88888888, 0xdb3b4888, 0x88888888, 0xd3bb4888, 0x88888888, 0x33bbb488, 0x88888888, 0x43bbb488, 0x88888888, 0x843bb348, 0x88888888, 0x88439948, 0x88888888, 0x88849224, 0x88888888, 0x88842248, 0x88888888})
var characterMoving5, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving500, spriteCharacterMoving510, spriteCharacterMoving501, spriteCharacterMoving511, spriteCharacterMoving502, spriteCharacterMoving512})

var spriteCharacterMoving600, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving610, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving601, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x84888888, 0x88888888, 0x46484444, 0x88888888, 0x4644666a, 0x88888884, 0x222266aa, 0x88888884, 0x22266aaa, 0x88888884, 0x666aaaa6, 0x88888884, 0x2266aa00, 0x88888884, 0x666a6092, 0x88888884, 0x222660dd, 0x88888884, 0x2106d092, 0x88888888, 0x4112d011, 0x88888888, 0x422dd111, 0x88888888, 0x84d00111, 0x88888888, 0x844d0011, 0x88888888, 0x45c557da})
var spriteCharacterMoving611, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88444488, 0x88888888, 0x4466aa48, 0x88888888, 0xaaa6aaa4, 0x88888888, 0xaaaaaaa4, 0x88888888, 0xaaaa6aa4, 0x88888888, 0xaaaa6a48, 0x88888888, 0x6a66a488, 0x88888888, 0x00044888, 0x88888888, 0x92024888, 0x88888888, 0xdad48444, 0x88888888, 0x0da44eb4, 0x88888888, 0x1ad4eec4, 0x88888888, 0x1119eee4, 0x88888888, 0x00197e48, 0x88888888, 0x11299488, 0x88888888, 0xd9b99488, 0x88888888})
var spriteCharacterMoving602, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888884, 0xcc5c5577, 0x8888884c, 0xccc55ccc, 0x888884cc, 0xccc5cccc, 0x888884cc, 0xc5c5ccc6, 0x888884cc, 0x5c0116bd, 0x8888884c, 0x5c1111dd, 0x88888884, 0x421110bc, 0x88888888, 0x42911555, 0x88888888, 0x4999399d, 0x88888888, 0x4994bbbb, 0x88888888, 0x8444bddd, 0x88888888, 0x4339dddd, 0x88888888, 0x4339dddd, 0x88888888, 0x4344bdb4, 0x88888888, 0x44884448, 0x88888888, 0x88888888})
var spriteCharacterMoving612, _ = spritengine.CreateSprite(paletteCharacter, []int{0x96bb4888, 0x88888888, 0xc6db4888, 0x88888888, 0x6ddb4488, 0x88888888, 0xdddcd048, 0x88888888, 0xddcc0048, 0x88888888, 0xdcccd048, 0x88888888, 0xccc55488, 0x88888888, 0x55544888, 0x88888888, 0x00948888, 0x88888888, 0xbd348888, 0x88888888, 0xdd3b4888, 0x88888888, 0xd3bb4888, 0x88888888, 0x3bb34888, 0x88888888, 0x49948888, 0x88888888, 0x42244888, 0x88888888, 0x42224888, 0x88888888})
var characterMoving6, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving600, spriteCharacterMoving610, spriteCharacterMoving601, spriteCharacterMoving611, spriteCharacterMoving602, spriteCharacterMoving612})

var spriteCharacterMoving700, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving710, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving701, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88884444, 0x88888888, 0x44446666, 0x88888884, 0x622666aa, 0x88888884, 0x2266aaaa, 0x88888884, 0x66aaaa66, 0x88888884, 0x266aa000, 0x88888884, 0x66a60292, 0x88888884, 0x22660ddd, 0x88888884, 0x006d0290, 0x88888884, 0x112d0111, 0x88888888, 0x42dd1111, 0x88888888, 0x84001110, 0x88888888, 0x84400111, 0x88888884, 0x45c557da, 0x8888884c, 0xcc5c5577, 0x888884cc, 0xccc55ccc})
var spriteCharacterMoving711, _ = spritengine.CreateSprite(paletteCharacter, []int{0x44444888, 0x88888888, 0xaa6aa488, 0x88888888, 0xaa6aaa48, 0x88888888, 0xaaaaaa48, 0x88888888, 0xa666a488, 0x88888888, 0x00064888, 0x88888888, 0x90924888, 0x88888888, 0xdad48888, 0x88888888, 0xaa948888, 0x88888888, 0xda148444, 0x88888888, 0x11144eb4, 0x88888888, 0x0119eee4, 0x88888888, 0x11299e48, 0x88888888, 0xd9b99488, 0x88888888, 0x96bb4888, 0x88888888, 0xbddb4888, 0x88888888})
var spriteCharacterMoving702, _ = spritengine.CreateSprite(paletteCharacter, []int{0x888884cc, 0xccc5ccc6, 0x88888455, 0x5cc5cced, 0x888884cc, 0xccc566dd, 0x88888840, 0x1115bddd, 0x88888841, 0x1109bbbc, 0x88888884, 0x11195555, 0x88888888, 0x4999399d, 0x88888888, 0x4994bbbb, 0x88888888, 0x84443bdd, 0x88888888, 0x8884333d, 0x88888888, 0x88843b99, 0x88888888, 0x8884bb33, 0x88888888, 0x8843bb33, 0x88888888, 0x88499434, 0x88888888, 0x88422448, 0x88888888, 0x88422248})
var spriteCharacterMoving712, _ = spritengine.CreateSprite(paletteCharacter, []int{0xdddc4444, 0x88888888, 0xddcc4000, 0x48888888, 0xdccca000, 0x48888888, 0xccccddd0, 0x48888888, 0xccc55d00, 0x48888888, 0x55544444, 0x88888888, 0x00948888, 0x88888888, 0xddb48888, 0x88888888, 0xddd48888, 0x88888888, 0xddb48888, 0x88888888, 0xbd488888, 0x88888888, 0x94888888, 0x88888888, 0x48888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var characterMoving7, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving700, spriteCharacterMoving710, spriteCharacterMoving701, spriteCharacterMoving711, spriteCharacterMoving702, spriteCharacterMoving712})

var spriteCharacterMoving800, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888884, 0x48884444})
var spriteCharacterMoving810, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x44844448, 0x88888888, 0xaa4aaa48, 0x88888888})
var spriteCharacterMoving801, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888884, 0x6444666a, 0x88888888, 0x422666aa, 0x88888884, 0x2266aaaa, 0x88888884, 0x66aaaa66, 0x88888884, 0x266aa000, 0x88888884, 0x66a60969, 0x88888884, 0x22660ddd, 0x88888884, 0x106d0920, 0x88888884, 0x012d0111, 0x88888888, 0x42dd1111, 0x88888888, 0x84001110, 0x88888888, 0x44400111, 0x88888844, 0xc5c557aa, 0x888884cc, 0xcc5c5577, 0x88884ccc, 0xccc55ccc, 0x88884ccc, 0xcc55ccc6})
var spriteCharacterMoving811, _ = spritengine.CreateSprite(paletteCharacter, []int{0xaa6aaa48, 0x88888888, 0xaaaaaa48, 0x88888888, 0xaaa6a488, 0x88888888, 0xa6664888, 0x88888888, 0x00048888, 0x88888888, 0x20694888, 0x88888888, 0xadd48888, 0x88888888, 0xda248888, 0x88888888, 0xaa148444, 0x88888888, 0x11144ec4, 0x88888888, 0x0119eee4, 0x88888888, 0x1129ee48, 0x88888888, 0xa9b99484, 0x44888888, 0x76bb9440, 0x00488888, 0x6ddb44a0, 0x00488888, 0xdddc57dd, 0x00488888})
var spriteCharacterMoving802, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88884555, 0xcc55cc6d, 0x88884ccc, 0xc4e6ebdd, 0x88884011, 0x144bbddd, 0x88884111, 0x0429bbcc, 0x88888411, 0x12995555, 0x88888844, 0x49993399, 0x88888888, 0x4994bbbb, 0x88888888, 0x844433bd, 0x88888888, 0x44443333, 0x88888884, 0x229bbbb4, 0x88888884, 0x2293b348, 0x88888884, 0x24444488, 0x88888884, 0x48888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterMoving812, _ = spritengine.CreateSprite(paletteCharacter, []int{0xddcc575d, 0x00488888, 0xdccc5575, 0x54888888, 0xcccc5554, 0x48888888, 0xccc54448, 0x88888888, 0x55548888, 0x88888888, 0xd0948888, 0x88888888, 0xbdb48888, 0x88888888, 0xdddd4888, 0x88888888, 0xbdddd488, 0x88888888, 0x44bdd488, 0x88888888, 0x88499488, 0x88888888, 0x88433488, 0x88888888, 0x88433348, 0x88888888, 0x88444448, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var characterMoving8, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterMoving800, spriteCharacterMoving810, spriteCharacterMoving801, spriteCharacterMoving811, spriteCharacterMoving802, spriteCharacterMoving812})

var spriteCharacterStanding00, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x44888444})
var spriteCharacterStanding10, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x84484448, 0x88888888, 0x4aa4aaa4, 0x88888888})
var spriteCharacterStanding01, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x46444666, 0x88888888, 0x8422666a, 0x88888888, 0x42266aaa, 0x88888888, 0x466aaaa6, 0x88888888, 0x4266aa00, 0x88888888, 0x466a6099, 0x88888888, 0x422660dd, 0x88888888, 0x4106d092, 0x88888888, 0x4112d011, 0x88888888, 0x842dd111, 0x88888888, 0x88400111, 0x88888888, 0x84440011, 0x88888888, 0x45c5557a, 0x88888884, 0xcc5c5557, 0x8888884c, 0xccc555cc, 0x8888884c, 0xccc55ccc})
var spriteCharacterStanding11, _ = spritengine.CreateSprite(paletteCharacter, []int{0xaaa6aaa4, 0x88888888, 0xaaaaaaa4, 0x88888888, 0xaaaa6a48, 0x88888888, 0x6a666488, 0x88888888, 0x00004888, 0x88888888, 0x99099488, 0x88888888, 0xdadd4888, 0x88888888, 0x0aa24884, 0x44888888, 0x1da1484e, 0x54888888, 0x111144ee, 0xb4888888, 0x0011293e, 0x48888888, 0x11129394, 0x88888888, 0xaa9b9934, 0x88888888, 0x79ebb548, 0x88888888, 0xc6ddb554, 0x88888888, 0x6dddc554, 0x88888888})
var spriteCharacterStanding02, _ = spritengine.CreateSprite(paletteCharacter, []int{0x888884cc, 0xccc55cc6, 0x888884cc, 0xcc45566d, 0x88888455, 0x5c44bbdd, 0x888884cc, 0xc4429bbb, 0x88888401, 0x14299555, 0x88888411, 0x11999399, 0x88888411, 0x10994bdb, 0x88888841, 0x14444ddd, 0x88888884, 0x4884dddd, 0x88888888, 0x8884dddd, 0x88888888, 0x884bdddd, 0x88888888, 0x884dddd4, 0x88888888, 0x884bdd44, 0x88888888, 0x88499488, 0x88888888, 0x88433488, 0x88888888, 0x88433488})
var spriteCharacterStanding12, _ = spritengine.CreateSprite(paletteCharacter, []int{0xdddcc555, 0x48888888, 0xddccc555, 0x48888888, 0xdcccc777, 0x48888888, 0xcccc5555, 0x48888888, 0x55554ddd, 0x48888888, 0xd1090d00, 0x48888888, 0xbb3300a0, 0x48888888, 0xdb334004, 0x88888888, 0xd3bb4448, 0x88888888, 0xb3bb4888, 0x88888888, 0x3bbb4888, 0x88888888, 0x3bb48888, 0x88888888, 0xbbb48888, 0x88888888, 0x49948888, 0x88888888, 0x42248888, 0x88888888, 0x42224888, 0x88888888})
var characterStanding, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterStanding00, spriteCharacterStanding10, spriteCharacterStanding01, spriteCharacterStanding11, spriteCharacterStanding02, spriteCharacterStanding12})

var spriteCharacterJumping00, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888844, 0x88888888, 0x84444466, 0x88888888, 0x46222666, 0x88888888, 0x422666aa})
var spriteCharacterJumping10, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x44444448, 0x88888888, 0x6aaa6aa4, 0x88888888, 0xaaaaa6aa, 0x48888888, 0xaaaaaaaa, 0x48888888})
var spriteCharacterJumping01, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x466aaaaa, 0x88888888, 0x42266aa0, 0x88888888, 0x4666a609, 0x88888888, 0x4222660d, 0x88888888, 0x42106d09, 0x88888888, 0x84112d01, 0x88888888, 0x8422dd11, 0x88888888, 0x884d0011, 0x88888888, 0x8844d001, 0x88888888, 0x845c5557, 0x88888888, 0x4cc5c555, 0x88888884, 0xccc5c5cc, 0x88888884, 0xcccc5ccc, 0x8888884c, 0xcc555ccc, 0x8888884c, 0xc5c1166b, 0x8888884c, 0xc51111bd})
var spriteCharacterJumping11, _ = spritengine.CreateSprite(paletteCharacter, []int{0x66a666aa, 0x48888888, 0x000066a4, 0x88888888, 0x99909448, 0x88888888, 0xddda4888, 0x88888888, 0x20aa4888, 0x88888888, 0x11ad4844, 0x48888888, 0x111144e5, 0x48888888, 0x10014eeb, 0x48888888, 0x111293e4, 0x88888888, 0xaa9b9948, 0x88888888, 0x797bb488, 0x88888888, 0xcc3dd488, 0x88888888, 0xc6ddd548, 0x88888888, 0x6dddc774, 0x88888888, 0xdddcc554, 0x88888888, 0xddcccddd, 0x48888888})
var spriteCharacterJumping02, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888884, 0xcc0111bb, 0x88888888, 0x4441105c, 0x88888888, 0x84299399, 0x88888888, 0x84999bbb, 0x88888888, 0x84994ddd, 0x88888888, 0x88444ddd, 0x88888888, 0x8884bddd, 0x88888888, 0x8884dddd, 0x88888888, 0x8884dddb, 0x88888888, 0x8884ddd4, 0x88888888, 0x884bdd48, 0x88888888, 0x884ddb48, 0x88888888, 0x88499488, 0x88888888, 0x88422488, 0x88888888, 0x88422488, 0x88888888, 0x88844888})
var spriteCharacterJumping12, _ = spritengine.CreateSprite(paletteCharacter, []int{0xcccc5d00, 0x48888888, 0xccc500d0, 0x48888888, 0xd1094004, 0x88888888, 0xba3b4448, 0x88888888, 0xdb3bb488, 0x88888888, 0xd3bbb488, 0x88888888, 0xb3bbb488, 0x88888888, 0x39bb3488, 0x88888888, 0x42994888, 0x88888888, 0x42248888, 0x88888888, 0x42248888, 0x88888888, 0x84488888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})

var characterJumping, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterJumping00, spriteCharacterJumping10, spriteCharacterJumping01, spriteCharacterJumping11, spriteCharacterJumping02, spriteCharacterJumping12})

var spriteCharacterLanding00, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x84888884, 0x88888888, 0x4648444a, 0x88888888, 0x4644666a, 0x88888888, 0x422266aa})
var spriteCharacterLanding10, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x84444488, 0x88888888, 0x4666aa48, 0x88888888, 0xaaa6aaa4, 0x88888888, 0xaaaaaaa4, 0x88888888, 0xaaaa6a48, 0x88888888})
var spriteCharacterLanding01, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888888, 0x42266aaa, 0x88888888, 0x466aaaa6, 0x88888888, 0x4266aa00, 0x88888888, 0x466a6099, 0x88888888, 0x422660dd, 0x88888888, 0x4106d092, 0x88888888, 0x4112d011, 0x88888888, 0x842dd111, 0x88888884, 0x44400111, 0x8888884c, 0xcc5c5011, 0x888884cc, 0xccc5c577, 0x888884c5, 0x55c5c555, 0x8888845c, 0xcc5c5ccc, 0x888884c0, 0x11445cc5, 0x88888451, 0x11145666, 0x88888841, 0x1104bbbd})
var spriteCharacterLanding11, _ = spritengine.CreateSprite(paletteCharacter, []int{0xaaaa6a48, 0x88888888, 0x6a66a488, 0x88888888, 0x00004888, 0x88888888, 0x99099488, 0x88888888, 0xdadd4884, 0x44888888, 0x0aa9484e, 0x54888888, 0x1da144ee, 0xb4888888, 0x1111293e, 0x48888888, 0x00119394, 0x88888888, 0x11129934, 0x88888888, 0xaa9bb544, 0x88888888, 0x72edd557, 0x48888888, 0xc3ddd575, 0x54888888, 0x6dddc75d, 0x00488888, 0xdddcc70a, 0xd0488888, 0xddcc5400, 0x0d488888})
var spriteCharacterLanding02, _ = spritengine.CreateSprite(paletteCharacter, []int{0x88888884, 0x11426bbb, 0x88888888, 0x4429955c, 0x88888888, 0x84999775, 0x88888888, 0x849943dd, 0x88888888, 0x88444bdd, 0x88888888, 0x88884ddd, 0x88888888, 0x88884ddd, 0x88888888, 0x88884ddd, 0x88888888, 0x88884bdd, 0x88888888, 0x888884b9, 0x88888888, 0x88888843, 0x88888888, 0x88888843, 0x88888888, 0x88888884, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888, 0x88888888})
var spriteCharacterLanding12, _ = spritengine.CreateSprite(paletteCharacter, []int{0xcccc4840, 0x04888888, 0xccc54884, 0x48888888, 0x55574888, 0x88888888, 0xbd094888, 0x88888888, 0xdb334888, 0x88888888, 0xdd3b3488, 0x88888888, 0xddbbb488, 0x88888888, 0xddbbb488, 0x88888888, 0xdb3bb348, 0x88888888, 0x943bbb48, 0x88888888, 0x344bbb48, 0x88888888, 0x3443b348, 0x88888888, 0x48849948, 0x88888888, 0x88842248, 0x88888888, 0x88842224, 0x88888888, 0x88884444, 0x88888888})
var characterLanding, _ = spritengine.CreateSpriteGroup(2, 3, &[]*spritengine.Sprite{spriteCharacterLanding00, spriteCharacterLanding10, spriteCharacterLanding01, spriteCharacterLanding11, spriteCharacterLanding02, spriteCharacterLanding12})

// getCharacter gets the controllable character object
func getCharacter(xPos float64, yPos float64) *spritengine.GameObject {

	return &spritengine.GameObject{
		CurrentState: "standing",
		States: spritengine.GameObjectStates{
			"standing": spritengine.SpriteSeries{
				Sprites:         []spritengine.SpriteInterface{characterStanding},
				CyclesPerSecond: 1,
			},
			"moving": spritengine.SpriteSeries{
				Sprites:         []spritengine.SpriteInterface{characterMoving1, characterMoving2, characterMoving3, characterMoving4, characterMoving5, characterMoving6, characterMoving7, characterMoving8},
				CyclesPerSecond: 2,
			},
			"jumping": spritengine.SpriteSeries{
				Sprites:         []spritengine.SpriteInterface{characterJumping},
				CyclesPerSecond: 1,
			},
			"landing": spritengine.SpriteSeries{
				Sprites:         []spritengine.SpriteInterface{characterLanding},
				CyclesPerSecond: 1,
			},
		},
		Position: spritengine.Vector{
			X: xPos,
			Y: yPos,
		},
		Mass: 0.4,
		Velocity: spritengine.Vector{
			X: 2,
			Y: 0,
		},
		Direction:        spritengine.DirStationary,
		IsFlipped:        false,
		IsControllable:   true,
		IsFloor:          false,
		IsInteractive:    true,
		IsHidden:         false,
		DynamicData:      spritengine.DynamicData{},
		FloorY:           0,
		EventHandler:     characterEventHandler,
		CollisionHandler: characterCollisionHandler,
	}

}

// characterCollisionHandler handles collision events for the character
func characterCollisionHandler(gameObject *spritengine.GameObject, collision spritengine.Collision) {

	if collision.GameObject.GetDynamicData("type", "") == "powerup" {
		gameObject.Velocity.X = 3
		gameObject.Level.Gravity = 0.3
		collision.GameObject.IsHidden = true
		collision.GameObject.IsInteractive = false
	}

}

// characterEventHandler handles events for the character sprite
func characterEventHandler(eventCode int, gameObject *spritengine.GameObject) {

	switch eventCode {

	case spritengine.EventFreefall:
		gameObject.CurrentState = "landing"

	case spritengine.EventFloorCollision:

		if gameObject.Direction == spritengine.DirStationary {
			gameObject.CurrentState = "standing"
		} else {
			gameObject.CurrentState = "moving"
		}

	case spritengine.EventDropOffLevel:
		gameObject.Position.X = 20
		gameObject.Position.Y = 100
		gameObject.IsFlipped = false

	}

}

// Sprite information for the floor
var paletteFloor = &spritengine.Palette{"5": color.RGBA{162, 199, 88, 255}, "0": color.RGBA{51, 101, 71, 255}, "1": color.RGBA{24, 44, 59, 255}, "2": color.RGBA{32, 64, 66, 255}, "3": color.RGBA{0, 0, 0, 0}, "4": color.RGBA{87, 153, 69, 255}}
var spriteFloor, _ = spritengine.CreateSprite(paletteFloor, []int{0x43343333, 0x43343333, 0x43344334, 0x43344334, 0x55455544, 0x54555544, 0x55555555, 0x55555555, 0x55555555, 0x55555555, 0x55445555, 0x54555445, 0x54555555, 0x45555554, 0x44554455, 0x45540554, 0x04544445, 0x05440050, 0x00440040, 0x00040400, 0x40400000, 0x44004404, 0x44004440, 0x44404404, 0x44044400, 0x04404000, 0x04044000, 0x00400000, 0x11140010, 0x01001000, 0x22210120, 0x12012101})
var floor, _ = spritengine.CreateSpriteGroup(1, 1, &[]*spritengine.Sprite{spriteFloor})

// getFloor gets a new floor object
func getFloor(xPos float64, yPos float64) *spritengine.GameObject {

	return &spritengine.GameObject{
		CurrentState: "default",
		States: spritengine.GameObjectStates{
			"default": spritengine.SpriteSeries{
				Sprites:         []spritengine.SpriteInterface{floor},
				CyclesPerSecond: 1,
			},
		},
		Position: spritengine.Vector{
			X: xPos,
			Y: yPos,
		},
		Mass: 0,
		Velocity: spritengine.Vector{
			X: 0,
			Y: 0,
		},
		Direction:        spritengine.DirStationary,
		IsFlipped:        false,
		IsControllable:   false,
		IsFloor:          true,
		IsInteractive:    true,
		IsHidden:         false,
		DynamicData:      spritengine.DynamicData{},
		FloorY:           0,
		EventHandler:     func(eventCode int, gameObject *spritengine.GameObject) {},
		CollisionHandler: func(gameObject *spritengine.GameObject, collision spritengine.Collision) {},
	}

}

// Sprite information for clouds
var paletteCloud = &spritengine.Palette{"0": color.RGBA{0, 0, 0, 0}, "1": color.RGBA{255, 255, 255, 255}}
var spriteCloud00, _ = spritengine.CreateSprite(paletteCloud, []int{0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000011, 0x00000000, 0x00000011, 0x00000000, 0x00000111, 0x00000000, 0x00001111, 0x00000000, 0x00011111, 0x00000000, 0x00011111, 0x00000000, 0x00011111, 0x00000000, 0x00011111})
var spriteCloud10, _ = spritengine.CreateSprite(paletteCloud, []int{0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00011111, 0x11111000, 0x01111111, 0x11111110, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111})
var spriteCloud20, _ = spritengine.CreateSprite(paletteCloud, []int{0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x10000000, 0x00000000, 0x11000000, 0x00000000, 0x11100000, 0x00000000, 0x11110000, 0x00000000, 0x11110000, 0x00000000, 0x11111000, 0x00000000, 0x11111000, 0x00000000, 0x11111000, 0x00000000})
var spriteCloud01, _ = spritengine.CreateSprite(paletteCloud, []int{0x00000000, 0x00011111, 0x00000001, 0x11111111, 0x00000011, 0x11111111, 0x00001111, 0x11111111, 0x00011111, 0x11111111, 0x00111111, 0x11111111, 0x00111111, 0x11111111, 0x01111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111})
var spriteCloud11, _ = spritengine.CreateSprite(paletteCloud, []int{0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111})
var spriteCloud21, _ = spritengine.CreateSprite(paletteCloud, []int{0x11111000, 0x00000000, 0x11111111, 0x00000000, 0x11111111, 0x11000000, 0x11111111, 0x11110000, 0x11111111, 0x11111000, 0x11111111, 0x11111100, 0x11111111, 0x11111110, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111})
var spriteCloud02, _ = spritengine.CreateSprite(paletteCloud, []int{0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x01111111, 0x11111111, 0x00111111, 0x11111111, 0x00011111, 0x11111111, 0x00001111, 0x11111111, 0x00000111, 0x11111111, 0x00000011, 0x11111111, 0x00000000, 0x01111110, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000})
var spriteCloud12, _ = spritengine.CreateSprite(paletteCloud, []int{0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x01111111, 0x11111110, 0x00111111, 0x11111000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000})
var spriteCloud22, _ = spritengine.CreateSprite(paletteCloud, []int{0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111111, 0x11111110, 0x11111111, 0x11111100, 0x11111111, 0x11111000, 0x11111111, 0x11110000, 0x11111111, 0x11100000, 0x11111111, 0x10000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000})
var cloud, _ = spritengine.CreateSpriteGroup(3, 3, &[]*spritengine.Sprite{spriteCloud00, spriteCloud10, spriteCloud20, spriteCloud01, spriteCloud11, spriteCloud21, spriteCloud02, spriteCloud12, spriteCloud22})

// getCloud gets a new cloud object
func getCloud(xPos float64, yPos float64) *spritengine.GameObject {

	return &spritengine.GameObject{
		CurrentState: "default",
		States: spritengine.GameObjectStates{
			"default": spritengine.SpriteSeries{
				Sprites:         []spritengine.SpriteInterface{cloud},
				CyclesPerSecond: 1,
			},
		},
		Position: spritengine.Vector{
			X: xPos,
			Y: yPos,
		},
		Mass: 0,
		Velocity: spritengine.Vector{
			X: 0,
			Y: 0,
		},
		Direction:        spritengine.DirStationary,
		IsFlipped:        false,
		IsControllable:   false,
		IsFloor:          false,
		IsInteractive:    false,
		IsHidden:         false,
		DynamicData:      spritengine.DynamicData{},
		FloorY:           0,
		EventHandler:     func(eventCode int, gameObject *spritengine.GameObject) {},
		CollisionHandler: func(gameObject *spritengine.GameObject, collision spritengine.Collision) {},
	}

}

// Sprite information for powerups
var palettePowerup = &spritengine.Palette{"b": color.RGBA{255, 208, 106, 255}, "c": color.RGBA{245, 209, 127, 255}, "0": color.RGBA{255, 255, 255, 255}, "7": color.RGBA{255, 203, 91, 255}, "9": color.RGBA{34, 30, 32, 255}, "3": color.RGBA{233, 182, 76, 255}, "a": color.RGBA{234, 194, 106, 255}, "e": color.RGBA{171, 132, 51, 255}, "1": color.RGBA{255, 251, 243, 255}, "2": color.RGBA{237, 180, 59, 255}, "d": color.RGBA{0, 0, 0, 0}, "8": color.RGBA{255, 226, 162, 255}, "4": color.RGBA{255, 238, 203, 255}, "5": color.RGBA{255, 248, 234, 255}, "6": color.RGBA{47, 40, 34, 255}}
var spritePowerup, _ = spritengine.CreateSprite(palettePowerup, []int{0xd9666666, 0x6666669d, 0x9e223acc, 0xcaa322e9, 0x62777778, 0x87777726, 0x627777b5, 0x1b777726, 0x63777780, 0x08777726, 0x637b8800, 0x014cb776, 0x6a400000, 0x000005a6, 0x6ab50000, 0x00001bc6, 0x6a7b5000, 0x0001b7a6, 0x6a778000, 0x000877c6, 0x6a778000, 0x00087736, 0x63774001, 0x10047736, 0x6277858b, 0x78447726, 0x6277c777, 0x777c7726, 0x9e223aac, 0xcae322e9, 0xd9666666, 0x6666669d})
var powerup, _ = spritengine.CreateSpriteGroup(1, 1, &[]*spritengine.Sprite{spritePowerup})

// getPowerup gets a new powerup object
func getPowerup(xPos float64, yPos float64) *spritengine.GameObject {

	return &spritengine.GameObject{
		CurrentState: "default",
		States: spritengine.GameObjectStates{
			"default": spritengine.SpriteSeries{
				Sprites:         []spritengine.SpriteInterface{powerup},
				CyclesPerSecond: 1,
			},
		},
		Position: spritengine.Vector{
			X: xPos,
			Y: yPos,
		},
		Mass: 0,
		Velocity: spritengine.Vector{
			X: 0,
			Y: 0,
		},
		Direction:        spritengine.DirStationary,
		IsFlipped:        false,
		IsControllable:   false,
		IsFloor:          false,
		IsInteractive:    true,
		IsHidden:         false,
		DynamicData:      spritengine.DynamicData{"type": "powerup"},
		FloorY:           0,
		EventHandler:     func(eventCode int, gameObject *spritengine.GameObject) {},
		CollisionHandler: func(gameObject *spritengine.GameObject, collision spritengine.Collision) {},
	}

}
