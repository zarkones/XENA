package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func ToolsInspector() fyne.Container {
	var inspector = container.NewVBox()

	return *inspector
}
