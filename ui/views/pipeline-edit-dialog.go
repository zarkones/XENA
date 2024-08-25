package views

import (
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func PipelineEditDialog(pipeline xenaC2.Pipeline, callback func()) {
	w := core.App.NewWindow("Edit Pipeline: " + pipeline.Name)
	w.Resize(fyne.NewSize(300, 0))
	w.CenterOnScreen()

	pipeName := widget.NewEntry()
	pipeName.SetPlaceHolder("Pipeline Name...")
	pipeName.SetText(pipeline.Name)
	pipeName.OnChanged = func(s string) {
		pipeline.Name = s
	}

	pipeDesc := widget.NewEntry()
	pipeDesc.SetPlaceHolder("Pipeline Description...")
	pipeDesc.SetText(pipeline.Desc)
	pipeDesc.OnChanged = func(s string) {
		pipeline.Desc = s
	}

	w.SetContent(container.NewVBox(
		widget.NewLabel("Edit the pipeline's metadata:"),
		pipeName,
		pipeDesc,
		container.NewHBox(
			widget.NewButton("CANCEL", func() {
				w.Close()
			}),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("CONFIRM", theme.ConfirmIcon(), func() {
				if err := xenaC2.UpsertPipeline(pipeline); err != nil {
					Warn("Failed to edit the pipeline, exception: " + err.Error())
				}
				callback()
				w.Close()
			}),
		),
	))

	w.Show()
}
