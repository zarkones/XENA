package views

import (
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AgentReverseDNSLookup(callback func(host string)) {
	w := core.App.NewWindow("XENA: Reverse DNS Lookup")
	w.Resize(fyne.NewSize(300, 0))
	w.CenterOnScreen()

	targetHostInput := widget.NewEntry()
	targetHostInput.SetText("1.1.1.1")

	runBtn := widget.NewButton("Run", func() {
		callback(targetHostInput.Text)
		w.Close()
	})

	w.SetContent(container.NewVBox(
		targetHostInput,
		runBtn,
	))

	w.Show()
}
