//+build js

package oak

import "github.com/oakmound/shiny/screen"

var (
	drawLoopPublishDef = func(tx screen.Texture) {
		windowControl.Upload(zeroPoint, winBuffer, winBuffer.Bounds())
	}
)
