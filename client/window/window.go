package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Window struct {
	FyneWindow *fyne.Window
}

func (self Window) NewFromFyneWindow(fyneWindow *fyne.Window) *Window {
	return &Window{
		FyneWindow: fyneWindow,
	}
}

func (self *Window) ShowAndRun() {
	(*self.FyneWindow).ShowAndRun()
}

func (self *Window) Content() fyne.CanvasObject {
	return (*self.FyneWindow).Content()
}

func (self *Window) SetContent(content fyne.CanvasObject) {
	(*self.FyneWindow).SetContent(content)
}

func (self *Window) Focus(element fyne.Focusable) {
	(*self.FyneWindow).Canvas().Focus(element)
}

func (self *Window) Popup() {
	popup := dialog.NewInformation("title", "message", *self.FyneWindow)
	popup.Show()
}
