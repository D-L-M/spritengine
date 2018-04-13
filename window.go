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

				// Throttle to the desired FPS
				if frameAgeNano < targetFrameAgeNano {
					time.Sleep(time.Duration(targetFrameAgeNano-frameAgeNano) * time.Nanosecond)
					frameAgeNano = targetFrameAgeNano
				}

				game.CurrentFrame++

				if game.CurrentFrame > game.TargetFrameRate {
					game.CurrentFrame = 1
				}

				frameAgeSeconds := (float64(frameAgeNano) / float64(1000000000))
				currentFrameRate := 1 / frameAgeSeconds

				// Repaint the stage
				lastPaintTimeNano = time.Now().UnixNano()
				stage := image.NewRGBA(image.Rect(0, 0, game.Width, game.Height))

				game.CurrentLevel().BeforePaint(game.CurrentLevel())
				game.CurrentLevel().Repaint(stage)
				game.FramePainter(stage, game.CurrentLevel(), currentFrameRate)
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
