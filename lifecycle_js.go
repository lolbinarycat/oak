//+build js

package oak

import (
	"image"
	"runtime"

	"github.com/oakmound/oak/v2/collision"
	"github.com/oakmound/oak/v2/dlog"
	"github.com/oakmound/oak/v2/event"
	omouse "github.com/oakmound/oak/v2/mouse"
	"github.com/oakmound/oak/v2/render"
	"github.com/oakmound/oak/v2/timing"
	"golang.org/x/exp/shiny/screen"
)

var (
	winBuffer     screen.Buffer
	screenControl screen.Screen
	windowControl screen.Window
)

func lifecycleLoop(inScreen screen.Screen) {
	dlog.Info("Init Lifecycle")

	firstScene = firstSceneJs

	screenControl = inScreen
	var err error

	// The window buffer represents the subsection of the world which is available to
	// be shown in a window.
	dlog.Info("Creating window buffer")
	winBuffer, err = screenControl.NewBuffer(image.Point{ScreenWidth, ScreenHeight})
	if err != nil {
		dlog.Error(err)
		return
	}

	// The window controller handles incoming hardware or platform events and
	// publishes image data to the screen.\
	dlog.Info("Creating window controller")
	changeWindow(ScreenWidth*conf.Screen.Scale, ScreenHeight*conf.Screen.Scale)

	dlog.Info("Getting event bus")
	eb = event.GetBus()

	go drawLoop()
	inputLoopInit()
	var prevScene string

	s, ok := sceneMap[firstScene]
	if !ok {
		dlog.Error("Unknown scene", firstScene)
		panic("Unknown scene")
	}
	s.active = true
	globalFirstScene = firstScene
	CurrentScene = "loading"

	result := new(SceneResult)

	dlog.Info("First Scene Start")

	drawCh <- true
	drawCh <- true

	dlog.Verb("Draw Channel Activated")

	schedCt := 0
	for {
		ViewPos = image.Point{0, 0}
		updateScreen(0, 0)
		useViewBounds = false

		dlog.Info("Scene Start", CurrentScene)
		go func() {
			dlog.Info("Starting scene in goroutine", CurrentScene)
			s, ok := sceneMap[CurrentScene]
			if !ok {
				dlog.Error("Unknown scene", CurrentScene)
				panic("Unknown scene")
			}
			s.start(prevScene, result.NextSceneInput)
			transitionCh <- true
		}()
		sceneTransition(result)
		// Post transition, begin loading animation
		dlog.Info("Starting load animation")
		drawCh <- true
		dlog.Info("Getting Transition Signal")
		<-transitionCh
		dlog.Info("Resume Drawing")
		// Send a signal to resume (or begin) drawing
		drawCh <- true

		dlog.Info("Looping Scene")
		cont := true
		logicTicker := logicTickerInit()
		for cont {
			logicLoopSingle(logicTicker)
			inputLoopSwitch()
			event.ResolvePendingSingle()
			cont = sceneMap[CurrentScene].loop()
			schedCt++
			if schedCt > 100 {
				schedCt = 0
				runtime.Gosched()
			}
		}
		dlog.Info("Scene End", CurrentScene)

		prevScene = CurrentScene

		// Send a signal to stop drawing
		drawCh <- true

		// Reset any ongoing delays
	delayLabel:
		for {
			select {
			case timing.ClearDelayCh <- true:
			default:
				break delayLabel
			}
		}

		dlog.Verb("Resetting Engine")
		// Reset transient portions of the engine
		// We start by clearing the event bus to
		// remove most ongoing code
		event.ResetBus()
		// We follow by clearing collision areas
		// because otherwise collision function calls
		// on non-entities (i.e. particles) can still
		// be triggered and attempt to access an entity
		dlog.Verb("Event Bus Reset")
		collision.Clear()
		omouse.Clear()
		event.ResetEntities()
		render.ResetDrawStack()
		render.PreDraw()
		dlog.Verb("Engine Reset")

		// Todo: Add in customizable loading scene between regular scenes

		CurrentScene, result = sceneMap[CurrentScene].end()
		// For convenience, we allow the user to return nil
		// but it gets translated to an empty result
		if result == nil {
			result = new(SceneResult)
		}

		eb = event.GetBus()
	}
}
