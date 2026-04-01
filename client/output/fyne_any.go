package output

import "fyne.io/fyne/v2"

func NewOutputFyneAny() (Output, fyne.CanvasObject) {
	return NewOutputFyneTextGrid() // NewOutputFyneRichText()
}
