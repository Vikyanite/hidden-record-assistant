package theme

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

var (
	//go:embed font/NotoSansSC-Regular.ttf
	FontBi []byte
)

type MyTheme struct{}

func (m *MyTheme) Font(fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "NotoSansSC-Regular.ttf",
		StaticContent: FontBi,
	}
}

func (*MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
