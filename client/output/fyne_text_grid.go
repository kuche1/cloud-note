package output

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type OutputFyneTextGrid struct {
	textGrid *widget.TextGrid
}

// Deprecated: (At least on desktop) when a message gets too long this does not wrap
// the text
func NewOutputFyneTextGrid() (*OutputFyneTextGrid, *widget.TextGrid) {
	widget := widget.NewTextGrid()
	// widget.Scroll = fyne.ScrollBoth // This isn't a great solution

	return &OutputFyneTextGrid{
		textGrid: widget,
	}, widget
}

func (self *OutputFyneTextGrid) Println(text string) {
	// I'm adding `fyne.Do` since as of right now everywhere in the code
	// where `.Append` is called is wrapped in `fyne.Do`
	fyne.Do(func() {
		self.textGrid.Append(text)
	})
}
