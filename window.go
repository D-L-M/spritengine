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

// CreateWindow creates a window and provides a corresponding image that can be drawn on
func CreateWindow(title string, width int, height int, scaleFactor int, targetFrameRate int, framePainter FramePainter, keyListener KeyListener) {

	lastPaintTimeNano := time.Now().UnixNano()
	targetFrameAgeNano := int64(1000000000) / int64(targetFrameRate)

	driver.Main(func(src screen.Screen) {

		res := image.Pt(width*scaleFactor, height*scaleFactor)
		win, _ := src.NewWindow(&screen.NewWindowOptions{Width: res.X, Height: res.Y, Title: title})
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

				lastPaintTimeNano = timeNowNano
				frameAgeSeconds := (float64(frameAgeNano) / float64(1000000000))
				currentFrameRate := 1 / frameAgeSeconds

				stage := image.NewRGBA(image.Rect(0, 0, width, height))

				framePainter(stage, currentFrameRate)
				xdraw.NearestNeighbor.Scale(buf.RGBA(), image.Rect(0, 0, width*scaleFactor, height*scaleFactor), stage, stage.Bounds(), draw.Over, nil)
				win.Upload(image.Point{}, buf, buf.Bounds())
				win.Publish()

				win.Send(paint.Event{})

				// Key presses
			case key.Event:

				keyListener(e)

			}

		}

	})

}
