package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (self *Settings) PromptNewServerAddr(window *fyne.Window, info string) (_ok bool) {
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
					// Improve: This will block

					err := self.SetServerAddr(entry.Text)

					if err != nil {
						// Improve: This is terrible

						fyne.Do(func() {
							button12345 := widget.NewButton(
								"ok",
								func() { wasChanged <- false },
							)
							dialog12345 := dialog.NewCustomWithoutButtons(
								"Error",
								button12345,
								*window,
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
			*window,
		)

		dialog.Show()
	})

	return <-wasChanged
}
