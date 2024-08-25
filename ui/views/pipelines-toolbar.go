package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func PipelinesToolbar() fyne.CanvasObject {
	toolbar := container.NewHBox(
		widget.NewButtonWithIcon("New Pipeline", theme.ContentAddIcon(), func() {
			NewPipeline()
		}),
	)

	return toolbar
}
