package views

import (
	"ui/core"
	"ui/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func NewPipeline() {
	w := core.App.NewWindow("XENA: New Pipeline")
	w.Resize(fyne.NewSize(640, 0))
	w.CenterOnScreen()

	pipelineNameInput := widget.NewEntry()
	pipelineNameInput.SetPlaceHolder("Name")

	pipelineDescInput := widget.NewMultiLineEntry()
	pipelineDescInput.SetPlaceHolder("Description")

	w.SetContent(container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Create New Pipeline"),
		widget.NewLabel("Pipeline is a series of steps executing a tool or a script, allowing you to automate tasks and create custom scanners."),
		pipelineNameInput,
		pipelineDescInput,

		widget.NewButton("CREATE", func() {
			newPipeline := xenaC2.Pipeline{
				Name: pipelineNameInput.Text,
				Desc: pipelineDescInput.Text,
			}

			if err := xenaC2.UpsertPipeline(newPipeline); err != nil {
				Warn("Failed to Create the New Pipeline, exception: " + err.Error())
				return
			}

			state.Pipelines, _ /*err*/ = xenaC2.GetPipelines()
			refreshPipelinesTable()

			w.Close()
		}),
	))

	w.Show()
}
