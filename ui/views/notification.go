package views

import (
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const notifyWindowWidth = 160
const notifyWindowHeight = 80

func Notify(title, msg string) {
	w := core.App.NewWindow(title)
	w.Resize(fyne.NewSize(notifyWindowWidth, notifyWindowHeight))
	w.SetContent(container.New(layout.NewHBoxLayout(), widget.NewLabel(msg)))
	w.CenterOnScreen()
	w.Show()
}

func Into(msg string) {
	Notify("XENA: New Notification", msg)
}

func Warn(msg string) {
	Notify("XENA: New Warning!", msg)
}

func Alert(msg string) {
	Notify("XENA: New Alert!", msg)
}
