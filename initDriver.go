// +build !js

package oak

import "github.com/oakmound/oak/v2/dlog"

var firstSceneJs string

func initDriver(firstScene, imageDir, audioDir string) {
	dlog.Info("Init Scene Loop")
	go sceneLoop(firstScene, conf.TrackInputChanges)
	dlog.Info("Init Console")
	go defaultDebugConsole()
	dlog.Info("Init Main Driver")
	firstSceneJs = firstScene
	InitDriver(lifecycleLoop)
}
