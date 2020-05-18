//+build !js

package oak

import (
	"image"

	"github.com/oakmound/oak/v2/dlog"
	"golang.org/x/mobile/event/lifecycle"

	"github.com/oakmound/shiny/screen"
)

var (
	winBuffer     screen.Image
	screenControl screen.Screen
	windowControl screen.Window
)

func lifecycleLoop(s screen.Screen) {
	dlog.Info("Init Lifecycle")

	screenControl = s
	var err error

	// The window buffer represents the subsection of the world which is available to
	// be shown in a window.
	dlog.Info("Creating window buffer")
	winBuffer, err = screenControl.NewImage(image.Point{ScreenWidth, ScreenHeight})
	if err != nil {
		dlog.Error(err)
		return
	}

	// Next time:
	// Right here, query the backing scale factor of the physical screen
	// Apply that factor to the scale

	dlog.Info("Creating window controller")
	changeWindow(int32(conf.Screen.X), int32(conf.Screen.Y), ScreenWidth*conf.Screen.Scale, ScreenHeight*conf.Screen.Scale)

	dlog.Info("Starting draw loop")
	go drawLoop()
	dlog.Info("Starting input loop")
	go inputLoop()

	// The quit channel represents a signal
	// for the engine to stop.
	// Lifecycle goroutine reaches here.
	<-quitCh
}

// Quit sends a signal to the window to close itself, ending oak.
func Quit() {
	windowControl.Send(lifecycle.Event{To: lifecycle.StageDead})
}
