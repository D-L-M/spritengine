package spritengine

import (
	"image"
	"image/draw"
	"time"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	xdraw "golang.org/x/image/draw"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

// createWindow creates a window and provides a corresponding image that can be drawn on
func createWindow(game *Game) {

	lastPaintTimeNano := time.Now().UnixNano()
	targetFrameAgeNano := int64(1000000000) / int64(game.TargetFrameRate)
	previousFrameRates := []float64{}
	maxFrameRatesToConsider := 100
	frameRateDrift := 0.0

	// Prefill the previous framerate slice
	for i := 0; i < maxFrameRatesToConsider; i++ {
		previousFrameRates = append(previousFrameRates, float64(game.TargetFrameRate))
	}

	// Create a new window and listen for events
	driver.Main(func(src screen.Screen) {

		res := image.Pt(game.Width*game.ScaleFactor, game.Height*game.ScaleFactor)
		win, _ := src.NewWindow(&screen.NewWindowOptions{Width: res.X, Height: res.Y, Title: game.Title})
		buf, _ := src.NewBuffer(res)

		for {

			switch event := win.NextEvent().(type) {

			// Close the window
			case lifecycle.Event:

				if event.To == lifecycle.StageDead {
					return
				}

				// Window repaints
			case paint.Event:

				frameAgeNano := (time.Now().UnixNano() - lastPaintTimeNano)
				frameRateDriftNano := int64(frameRateDrift * 1000000000)

				// Throttle to the desired FPS
				if frameAgeNano < targetFrameAgeNano {
					time.Sleep(time.Duration(targetFrameAgeNano-frameAgeNano-frameRateDriftNano) * time.Nanosecond)
					frameAgeNano = targetFrameAgeNano
				}

				game.CurrentFrame++

				if game.CurrentFrame > game.TargetFrameRate {
					game.CurrentFrame = 1
				}

				frameAgeSeconds := (float64(frameAgeNano) / float64(1000000000))
				currentFrameRate := 1 / frameAgeSeconds

				// Work out the average frame rate
				previousFrameRates = append(previousFrameRates, currentFrameRate)

				if len(previousFrameRates) > int(maxFrameRatesToConsider) {
					previousFrameRates = previousFrameRates[1:]
				}

				talliedFrameRate := 0.0

				for _, individualFrameRate := range previousFrameRates {
					talliedFrameRate += individualFrameRate
				}

				averageFrameRate := talliedFrameRate / float64(maxFrameRatesToConsider)

				// Figure out by how much we're straying from the target framerate
				frameRateDrift = float64(game.TargetFrameRate) - averageFrameRate

				// Repaint the stage
				lastPaintTimeNano = time.Now().UnixNano()
				stage := image.NewRGBA(image.Rect(0, 0, game.Width, game.Height))

				game.CurrentLevel().BeforePaint(game.CurrentLevel())
				game.CurrentLevel().Repaint(stage)
				game.FramePainter(stage, game.CurrentLevel(), averageFrameRate)
				xdraw.NearestNeighbor.Scale(buf.RGBA(), image.Rect(0, 0, game.Width*game.ScaleFactor, game.Height*game.ScaleFactor), stage, stage.Bounds(), draw.Over, nil)
				win.Upload(image.Point{}, buf, buf.Bounds())
				win.Publish()

				win.Send(paint.Event{})

				// Key presses
			case key.Event:
				game.BroadcastInput(event)
			}

		}

	})

}
