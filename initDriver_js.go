//+build js

package oak

import "github.com/oakmound/oak/v2/dlog"

func initDriver(firstScene, imageDir, audioDir string) {
	dlog.Info("Init JS Driver")
	firstSceneJs = firstScene
	InitDriver(lifecycleLoop)
}
