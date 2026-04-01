package output

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type OutputFyneLabel struct {
	label *widget.Label
}

func NewOutputFyneLabel() (*OutputFyneLabel, *container.Scroll) {
	label := widget.NewLabel("")
	label.Wrapping = fyne.TextWrapWord
	label.Alignment = fyne.TextAlignCenter
	// label.Refresh() // not really needed in this scase

	scroll := container.NewVScroll(label)
	// actually, not having this is what causes the bug of making the app very wide vertically

	return &OutputFyneLabel{
		label: label,
	}, scroll
}

func (self *OutputFyneLabel) Println(text string) {
	fyne.Do(func() {
		self.label.Text += text
		self.label.Refresh()
	})
}
