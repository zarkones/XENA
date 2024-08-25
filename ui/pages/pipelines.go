package pages

import (
	"ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Pipelines() fyne.CanvasObject {
	return container.NewBorder(
		views.PipelinesToolbar(),
		nil,
		nil,
		nil,
		views.PipelinesTable(),
	)
}
