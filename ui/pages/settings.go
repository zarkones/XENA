package pages

import (
	"ui/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func Settings() fyne.CanvasObject {
	c2Host := widget.NewEntry()
	c2Host.Text = *xenaC2.BaseURL
	c2Host.OnChanged = func(s string) {
		state.C2Host = s
	}

	authToken := widget.NewEntry()
	authToken.Text = *xenaC2.AuthToken
	authToken.OnChanged = func(s string) {
		state.AuthToken = s
	}

	return container.NewBorder(
		// Top.
		nil,

		// Bottom.
		nil,

		// Left.
		nil,

		// Right.
		nil,

		container.NewVBox(
			widget.NewLabel("C2 Host:"),
			c2Host,
			widget.NewLabel("Authentication Token:"),
			authToken,
			layout.NewSpacer(),
			widget.NewLabel("XENA (Beta) v0.3.2"),
		),
	)
}
