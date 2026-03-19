package output

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type OutputFyneTextGrid struct {
	textGrid *widget.TextGrid
}

func NewOutputFyneTextGrid() (*OutputFyneTextGrid, *widget.TextGrid) {
	widget := widget.NewTextGrid()

	return &OutputFyneTextGrid{
		textGrid: widget,
	}, widget
}

func (self *OutputFyneTextGrid) Println(text string) {
	fyne.Do(func() {
		self.textGrid.Append(text)
	})
}
