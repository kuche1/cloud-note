package settings

import "github.com/kuche1/cloud-note/client/window"

func (self *Settings) SceneInputEssential(window *window.Window, callbackWhenAllDone func()) {
	self.sceneInputMissingServerAddr(window, callbackWhenAllDone)
}
