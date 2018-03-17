package spritengine

import (
	"golang.org/x/mobile/event/key"
)

// Game is a struct that defines a game and the window that contains it
type Game struct {
	Title           string
	Width           int
	Height          int
	ScaleFactor     int
	TargetFrameRate int
	FramePainter    FramePainter
	KeyListener     KeyListener
	Levels          []*Level
	CurrentLevelID  int
	CurrentFrame    int
}

// CreateGame sets up a game and its window
func CreateGame(title string, width int, height int, scaleFactor int, targetFrameRate int, framePainter FramePainter, keyListener KeyListener, levels []*Level) *Game {

	game := Game{
		Title:           title,
		Width:           width,
		Height:          height,
		ScaleFactor:     scaleFactor,
		TargetFrameRate: targetFrameRate,
		FramePainter:    framePainter,
		KeyListener:     keyListener,
		Levels:          levels,
		CurrentLevelID:  0,
		CurrentFrame:    0,
	}

	for _, level := range levels {
		level.Game = &game
	}

	createWindow(&game)

	return &game

}

// CurrentLevel gets the current Level object
func (game *Game) CurrentLevel() *Level {

	return game.Levels[game.CurrentLevelID]

}

// BroadcastInput sends the game input to the current level's objects if they are controllable
func (game *Game) BroadcastInput(event key.Event) {

	for _, gameObject := range game.CurrentLevel().GameObjects {

		if gameObject.Controllable == true {
			game.KeyListener(event, gameObject)
		}

	}

}
