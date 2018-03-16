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
	Worlds          []*World
	CurrentWorldID  int
}

// CreateGame sets up a game and its window
func CreateGame(title string, width int, height int, scaleFactor int, targetFrameRate int, framePainter FramePainter, keyListener KeyListener, worlds []*World) *Game {

	game := Game{
		Title:           title,
		Width:           width,
		Height:          height,
		ScaleFactor:     scaleFactor,
		TargetFrameRate: targetFrameRate,
		FramePainter:    framePainter,
		KeyListener:     keyListener,
		Worlds:          worlds,
		CurrentWorldID:  0,
	}

	for _, world := range worlds {
		world.Game = &game
	}

	createWindow(&game)

	return &game

}

// CurrentWorld gets the current World object
func (game *Game) CurrentWorld() *World {

	return game.Worlds[game.CurrentWorldID]

}
