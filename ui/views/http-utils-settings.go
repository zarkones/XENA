package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func HttpUtilsSettings() fyne.CanvasObject {
	newListenerBtn := widget.NewButtonWithIcon("New Listener", theme.ContentAddIcon(), func() {
		// TODO: Add a HTTP listener on C2 server to connect apps like browser to.
	})

	heading := container.NewGridWithColumns(3, widget.NewLabel("Running"), widget.NewLabel("Interface"), container.NewHBox(layout.NewSpacer(), newListenerBtn))

	return container.NewVBox(
		heading,
	)
}
