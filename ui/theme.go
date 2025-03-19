package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type mainTheme struct{}

var COLOR_PRIMARY = color.RGBA{189, 147, 249, 255}
var COLOR_BG = color.RGBA{40, 42, 54, 255}
var COLOR_BG_2 = color.RGBA{68, 71, 90, 255}
var COLOR_ACTIVE = color.RGBA{68, 71, 90, 255}
var COLOR_RED = color.RGBA{255, 85, 85, 255}
var TRANSPARENT = color.RGBA{255, 255, 255, 0}

func (m mainTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground, theme.ColorNameInputBackground:
		return COLOR_BG

	case theme.ColorNameMenuBackground:
		return COLOR_ACTIVE

	case theme.ColorNameError:
		return COLOR_RED

	case theme.ColorNamePrimary, theme.ColorNameInputBorder, theme.ColorNameScrollBar:
		return COLOR_PRIMARY

	case theme.ColorNameSeparator, theme.ColorNameSelection, theme.ColorNameHover:
		return TRANSPARENT

	case theme.ColorNameButton:
		return COLOR_PRIMARY

	case theme.ColorNamePressed:
		return COLOR_BG

	case theme.ColorNameHeaderBackground:
		return COLOR_BG_2

	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (m mainTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m mainTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m mainTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case "innerPadding":
		return theme.DefaultTheme().Size(name) - theme.DefaultTheme().Size(name)/4

	case "inputRadius", "selectionRadius":
		return 0
	}

	return theme.DefaultTheme().Size(name) - theme.DefaultTheme().Size(name)*0.2
}

var _ fyne.Theme = (*mainTheme)(nil)
