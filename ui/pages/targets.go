package pages

import (
	"time"
	"ui/state"
	"ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func Targets() fyne.CanvasObject {
	var targets []xenaC2.Target

	targetsTable := widget.NewList(
		func() int {
			return len(targets)
		},

		func() fyne.CanvasObject {
			return widget.NewButton("mylongandlenghtydomain.com", func() {})
		},

		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Button).SetText(targets[i].Value)
			o.(*widget.Button).OnTapped = func() {
				// TODO
			}
		},
	)

	updateTargets := func() {
		targets, _ = xenaC2.GetTargets()
		targetsTable.Refresh()
	}

	go func() {
		go updateTargets()
		for range time.Tick(time.Second * 5) {
			if state.AuthToken == "" {
				continue
			}
			updateTargets()
		}
	}()

	addTargetBtn := widget.NewButtonWithIcon("Add Target", theme.ContentAddIcon(), func() {
		views.AddTarget(updateTargets)
	})

	return container.NewBorder(
		container.NewHBox(addTargetBtn),
		nil,
		targetsTable,
		nil,
	)
}
