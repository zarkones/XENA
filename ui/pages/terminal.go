package pages

import (
	"ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Terminal() fyne.CanvasObject {
	return container.NewBorder(
		views.AssetsToolbar(),
		nil,
		nil,
		nil,
		views.AgentsTable(),
	)
}
