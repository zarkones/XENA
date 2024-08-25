package pages

import (
	"ui/data"
	"ui/effects"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Shop() fyne.CanvasObject {

	productsList := container.NewVBox(
		effects.Gradient(
			widget.NewCard(
				"XENA: Mobile",
				"Best hacking toolkit running on Android.",
				container.NewHBox(
					layout.NewSpacer(),
					widget.NewHyperlink("Learn More / Download", data.XenaMobileURL),
					layout.NewSpacer(),
				),
			),
			false, true),
	)

	return container.NewBorder(
		// Top.
		productsList,
		// Bottom.
		nil,
		// Left.
		nil,
		// Right.
		nil,
	)
}
