package views

import (
	"encoding/json"
	"os"
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	xenaC2 "github.com/zarkones/xena-client"
)

func PipelineImportDialog(callback func(pipeline xenaC2.Pipeline)) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			Notify("Alert", err.Error())
			return
		}

		rawPipeline, err := os.ReadFile(reader.URI().Path())
		if err != nil {
			Notify("Alert", err.Error())
			return
		}

		var pipeline xenaC2.Pipeline

		if err := json.Unmarshal(rawPipeline, &pipeline); err != nil {
			Notify("Alert", err.Error())
			return
		}

		callback(pipeline)
	}, core.MainW)

	fd.Show()
}
