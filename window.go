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

	driver.Main(func(src screen.Screen) {

		res := image.Pt(game.Width*game.ScaleFactor, game.Height*game.ScaleFactor)
		win, _ := src.NewWindow(&screen.NewWindowOptions{Width: res.X, Height: res.Y, Title: game.Title})
		buf, _ := src.NewBuffer(res)

		for {

			switch e := win.NextEvent().(type) {

			// Close the window
			case lifecycle.Event:

				if e.To == lifecycle.StageDead {
					return
				}

				// Window repaints
			case paint.Event:

				timeNowNano := time.Now().UnixNano()
				frameAgeNano := (timeNowNano - lastPaintTimeNano)

				// Throttle to the desired FPS
				if frameAgeNano < targetFrameAgeNano {
					time.Sleep(time.Duration(targetFrameAgeNano-frameAgeNano) * time.Nanosecond)
					frameAgeNano = targetFrameAgeNano
				}

				// TODO: Move this code into the Game object?
				game.CurrentFrame++

				if game.CurrentFrame > game.TargetFrameRate {
					game.CurrentFrame = 1
				}

				lastPaintTimeNano = timeNowNano
				frameAgeSeconds := (float64(frameAgeNano) / float64(1000000000))
				currentFrameRate := 1 / frameAgeSeconds

				stage := image.NewRGBA(image.Rect(0, 0, game.Width, game.Height))

				game.CurrentLevel().Repaint(stage)
				game.FramePainter(stage, game.CurrentLevel(), currentFrameRate)
				xdraw.NearestNeighbor.Scale(buf.RGBA(), image.Rect(0, 0, game.Width*game.ScaleFactor, game.Height*game.ScaleFactor), stage, stage.Bounds(), draw.Over, nil)
				win.Upload(image.Point{}, buf, buf.Bounds())
				win.Publish()

				win.Send(paint.Event{})

				// Key presses
			case key.Event:

				game.BroadcastInput(e)
				game.KeyListener(e)

			}

		}

	})

}
