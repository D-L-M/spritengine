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
func CreateWindow(title string, width int, height int, scaleFactor int, framesPerSecond int, framePainter func(stage *image.RGBA), keyReaction func(key.Event)) {

	lastPaintTimeNano := time.Now().UnixNano()
	frameAgeNano := int64(1000000000) / int64(framesPerSecond)

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

				// Throttle to the desired FPS
				if (timeNowNano - lastPaintTimeNano) > frameAgeNano {

					lastPaintTimeNano = timeNowNano
					stage := image.NewRGBA(image.Rect(0, 0, width, height))

					framePainter(stage)
					xdraw.NearestNeighbor.Scale(buf.RGBA(), image.Rect(0, 0, width*scaleFactor, height*scaleFactor), stage, stage.Bounds(), draw.Over, nil)
					win.Upload(image.Point{}, buf, buf.Bounds())
					win.Publish()

				}

				win.Send(paint.Event{})

				// Key presses
			case key.Event:

				keyReaction(e)

			}

		}

	})

}
