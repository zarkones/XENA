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

func PipelineDeleteDialog(pipeline xenaC2.Pipeline, callback func()) {
	w := core.App.NewWindow("Delete Pipeline?")
	w.Resize(fyne.NewSize(300, 0))
	w.CenterOnScreen()

	w.SetContent(container.NewVBox(
		widget.NewLabel("Are you sure you want to delete the pipeline?"),
		container.NewHBox(
			widget.NewButton("CANCEL", func() {
				w.Close()
			}),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("DELETE", theme.DeleteIcon(), func() {
				if err := xenaC2.DeletePipeline(pipeline.ID); err != nil {
					Warn("Failed to delete the pipeline, exception: " + err.Error())
				}
				callback()
				w.Close()
			}),
		),
	))

	w.Show()
}
