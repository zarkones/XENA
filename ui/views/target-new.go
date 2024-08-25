package views

import (
	"ui/core"
	"ui/effects"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func AddTarget(callback func()) {
	w := core.App.NewWindow("XENA: Add Target")
	w.Resize(fyne.NewSize(300, 0))
	w.CenterOnScreen()

	target := xenaC2.Target{}

	typeSelect := widget.NewSelect([]string{"TEXT", "URL", "DOMAIN", "IP_ADDRESS_V4", "IP_ADDRESS_V6", "PHONE_NUMBER", "EMAIL_ADDRESS"}, func(value string) {
		target.Type = value
	})

	targetNameInput := widget.NewEntry()
	targetNameInput.SetPlaceHolder("Target Name")
	targetNameInput.OnChanged = func(s string) {
		target.Name = s
	}

	targetValueInput := widget.NewEntry()
	targetValueInput.SetPlaceHolder("Target Value, a domain, url, etc...")
	targetValueInput.OnChanged = func(s string) {
		target.Value = s
	}

	c := container.New(layout.NewVBoxLayout(),
		targetNameInput,

		targetValueInput,

		widget.NewLabel("Target Type"),
		typeSelect,

		widget.NewButton("CONFIRM", func() {
			if err := xenaC2.UpsertTargets(target); err != nil {
				Warn("Failed to save the Target, exception: " + err.Error())
				return
			}
			go callback()
			w.Close()
		}),
	)

	w.SetContent(effects.Gradient(c, true, true))

	w.Show()
}
