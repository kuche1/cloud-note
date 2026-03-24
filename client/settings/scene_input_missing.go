package settings

import "github.com/kuche1/cloud-note/client/window"

func (self *Settings) SceneInputMissing(window *window.Window, callbackWhenAllDone func()) {
	self.sceneInputMissingServerAddr(window, callbackWhenAllDone)
}
