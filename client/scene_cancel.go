package client

func (self *App) SceneCancel() {
	self.app.Quit() // TODO: Aparently this does nothing on mobile

	// self.FirstScene()
}
