package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const _PhoneScrollbarUnselectedSizeMultiplier = 1.8
const _PhoneScrollbarSelectedSizeMultiplier = 0.5

// ensure fyne.Theme is implemented
var _ fyne.Theme = (*Theme)(nil)

type Theme struct{}

func NewTheme() *Theme {
	return &Theme{}
}

func (t *Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t *Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *Theme) Size(name fyne.ThemeSizeName) float32 {
	if fyne.CurrentDevice().IsMobile() {
		if name == theme.SizeNameScrollBarSmall { // unselected (not being held down)
			return theme.DefaultTheme().Size(theme.SizeNameScrollBar) * _PhoneScrollbarUnselectedSizeMultiplier
		}
		if name == theme.SizeNameScrollBar { // selected (being held down)
			return theme.DefaultTheme().Size(theme.SizeNameScrollBar) * _PhoneScrollbarSelectedSizeMultiplier
		}
	}

	// if (name == theme.SizeNameScrollBar) || (name == theme.SizeNameScrollBarSmall) {
	// 	sizee := theme.DefaultTheme().Size(name)
	// 	return sizee * 3

	// 	// return 24
	// }

	return theme.DefaultTheme().Size(name)
}
