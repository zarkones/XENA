package views

import (
	"time"
	"ui/state"

	xenaC2 "github.com/zarkones/xena-client"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var pipelinesTable *widget.List

func refreshPipelinesTable() {
	pipelinesTable.Refresh()
}

func PipelinesTable() fyne.CanvasObject {
	updatePipelines := func() {
		// var err error
		state.Pipelines, _ /*err*/ = xenaC2.GetPipelines()
		// if err != nil {
		// 	Warn("failed to load pipelines:" + err.Error())
		// }
		refreshPipelinesTable()
	}

	pipelinesTable = widget.NewList(
		func() int {
			return len(state.Pipelines)
		},

		func() fyne.CanvasObject {
			return widget.NewCard("asdasdasdsad", "adasdasdasdasdasdaasdasdasdasdasd", container.NewHBox(widget.NewButton("Open in Editor", func() {})))
		},

		func(i widget.ListItemID, o fyne.CanvasObject) {
			pipe := state.Pipelines[i]
			o.(*widget.Card).SetTitle(pipe.Name)
			o.(*widget.Card).SetSubTitle(pipe.Desc)
			o.(*widget.Card).SetContent(container.NewVBox(
				container.NewHBox(
					layout.NewSpacer(),
					widget.NewButtonWithIcon("DELETE", theme.DeleteIcon(), func() {
						PipelineDeleteDialog(state.Pipelines[i], updatePipelines)
					}),
					widget.NewButtonWithIcon("EDIT", theme.DocumentCreateIcon(), func() {
						PipelineEditDialog(state.Pipelines[i], updatePipelines)
					}),
					widget.NewButtonWithIcon("OPEN IN EDITOR", theme.MediaPlayIcon(), func() {
						PipelineWindow(state.Pipelines[i])
					}),
				),
			))
		},
	)

	go func() {
		updatePipelines()
		for range time.Tick(time.Second * 5) {
			if state.AuthToken == "" {
				continue
			}
			updatePipelines()
		}
	}()

	return pipelinesTable
}
