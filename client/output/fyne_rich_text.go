package output

// import (
// 	"strings"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/widget"
// )

// type OutputFyneRichText struct {
// 	richText *widget.RichText
// }

// // TOD0: This is a mess, just comment out this shit
// //
// // IMPR0VE000: This is what causes the bug where after submitting a note the app increases
// // it's heights dramatically on desktop
// // Actually we can probably fix this issue just like we fixed the label issue my
// // adding a scroll, but I don't like the markdown stuff regardless
// func DeprecatedNewOutputFyneRichText() (*OutputFyneRichText, *widget.RichText) {
// 	widget := widget.NewRichText()
// 	widget.Wrapping = fyne.TextWrapWord

// 	return &OutputFyneRichText{
// 		richText: widget,
// 	}, widget
// }

// func (self *OutputFyneRichText) Println(text string) {
// 	// IMPR0VE000: This is not great, we need to find (or make) a better widget that
// 	// does not rely on hacks
// 	// IMPR0VE000: We can use this to insert 1 new line, but if we want to inser
// 	// 2 new lines this does not work
// 	fixedText := strings.ReplaceAll(text, "\n", "\n\n")

// 	fyne.Do(func() {
// 		// IMPR0VE000: Actually, does this automatically append a new line or not?
// 		// IMPR0VE000: "This API is intended for appending complete markdown documents or standalone fragments, and should not be used to parse a single markdown document piecewise"
// 		self.richText.AppendMarkdown(fixedText)
// 	})
// }
