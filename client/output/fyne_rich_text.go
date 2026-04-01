package output

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type OutputFyneRichText struct {
	richText *widget.RichText
}

// TODO: This is what causes the bug where after submitting a note the app increases
// it's heights dramatically on desktop
func DeprecatedNewOutputFyneRichText() (*OutputFyneRichText, *widget.RichText) {
	widget := widget.NewRichText()
	widget.Wrapping = fyne.TextWrapWord

	return &OutputFyneRichText{
		richText: widget,
	}, widget
}

func (self *OutputFyneRichText) Println(text string) {
	// IMPROVE000: This is not great, we need to find (or make) a better widget that
	// does not rely on hacks
	// IMPROVE000: We can use this to insert 1 new line, but if we want to inser
	// 2 new lines this does not work
	fixedText := strings.ReplaceAll(text, "\n", "\n\n")

	fyne.Do(func() {
		// IMPROVE000: Actually, does this automatically append a new line or not?
		// IMPROVE000: "This API is intended for appending complete markdown documents or standalone fragments, and should not be used to parse a single markdown document piecewise"
		self.richText.AppendMarkdown(fixedText)
	})
}
