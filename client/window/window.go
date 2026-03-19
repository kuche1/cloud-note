package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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

// TODO: Yes, just get rid of this
func (self *Window) ShowDialogOutput(
	title string,
	newThread func(output *widget.TextGrid),
) {
	fyne.Do(func() {
		output := widget.NewTextGrid()

		dialog := dialog.NewCustomWithoutButtons(
			title,
			output,
			*self.FyneWindow,
		)

		dialog.Show()

		go func() {
			newThread(output)

			fyne.Do(func() { dialog.Dismiss() })
			// Update:
			// IMPROVE000: I feel like this is going to bite me in the end
			//
			// Update:
			// Now that this is wrapped in `fyne.Do`, there will be less concern (yet still some)
			//
			// This may lead to some problems down the line, but it makes it impossible to accidentally forget to dismiss
			// a dialog, we'll see how it turns out
		}()
	})

}
