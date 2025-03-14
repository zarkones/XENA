package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func PipelinesToolbar() fyne.CanvasObject {
	toolbar := container.NewHBox(
		widget.NewButtonWithIcon("NEW", theme.ContentAddIcon(), func() {
			NewPipeline()
		}),
		widget.NewButtonWithIcon("IMPORT", theme.ContentAddIcon(), func() {
			PipelineImportDialog(func(pipeline xenaC2.Pipeline) {
				fmt.Println("pipeline", pipeline)
			})
		}),
	)

	return toolbar
}
