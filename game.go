package spritengine

// Game is a struct that defines a game and the window that contains it
type Game struct {
	Title           string
	Width           int
	Height          int
	ScaleFactor     int
	TargetFrameRate int
	FramePainter    FramePainter
	KeyListener     KeyListener
	CurrentWorld    *World
}

// CreateGame sets up a game and its window
func CreateGame(title string, width int, height int, scaleFactor int, targetFrameRate int, framePainter FramePainter, keyListener KeyListener, world *World) *Game {

	game := Game{
		Title:           title,
		Width:           width,
		Height:          height,
		ScaleFactor:     scaleFactor,
		TargetFrameRate: targetFrameRate,
		FramePainter:    framePainter,
		KeyListener:     keyListener,
		CurrentWorld:    world,
	}

	game.CurrentWorld.Game = &game

	createWindow(&game)

	return &game

}
