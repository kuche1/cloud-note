package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kuche1/cloud-note/client/window"
)

func (self *Settings) PromptNewServerAddr(window *window.Window, info string) (_ok bool) {
	wasChanged := make(chan bool)

	fyne.Do(func() {

		label := widget.NewLabel(info + "\n\nWould you like to set a new server address?")
		label.Wrapping = fyne.TextWrapWord // TextWrapBreak is really ugly here

		entry := widget.NewEntry()
		entry.Text = self.ServerAddr

		container := container.NewVBox(label, entry)

		dialog := dialog.NewCustomConfirm(
			"Change Server Address",
			"Set",
			"Cancel",
			container,
			func(yes bool) {
				if yes {
					// Improve000: This will block

					err := self.SetServerAddr(entry.Text)

					if err != nil {
						// Improve: This is terrible

						fyne.Do(func() {
							button12345 := widget.NewButton(
								"Ok",
								func() { wasChanged <- false },
							)
							dialog12345 := dialog.NewCustomWithoutButtons(
								"Error",
								button12345,
								*window.FyneWindow,
							)
							dialog12345.Show()
						})

						return
					}

					wasChanged <- true

				} else {
					wasChanged <- false
				}
			},
			*window.FyneWindow,
		)

		dialog.Show()
	})

	return <-wasChanged
}
