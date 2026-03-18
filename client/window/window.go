package window

import (
	"fyne.io/fyne/v2"
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

func (self *Window) SetContent(content fyne.CanvasObject) {
	(*self.FyneWindow).SetContent(content)
}
