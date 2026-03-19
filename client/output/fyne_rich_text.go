package output

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type OutputFyneRichText struct {
	richText *widget.RichText
}

func NewOutputFyneRichText() (*OutputFyneRichText, *widget.RichText) {
	widget := widget.NewRichText()
	widget.Wrapping = fyne.TextWrapWord

	return &OutputFyneRichText{
		richText: widget,
	}, widget
}

func (self *OutputFyneRichText) Println(text string) {
	fyne.Do(func() {
		self.richText.AppendMarkdown(text)
	})
}
