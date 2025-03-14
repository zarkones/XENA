package views

import (
	"encoding/json"
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	xenaC2 "github.com/zarkones/xena-client"
)

func PipelineExportDialog(pipeline xenaC2.Pipeline, callback func()) {
	serializedPipeline, err := json.Marshal(&pipeline)
	if err != nil {
		Notify("Alert", err.Error())
		return
	}

	fd := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		if uc == nil {
			return
		}
		if err != nil {
			Notify("Alert", err.Error())
			return
		}

		if _, err := uc.Write(serializedPipeline); err != nil {
			Notify("Alert", err.Error())
			return
		}

		uc.Close()
		callback()
	}, core.MainW)

	fd.Show()
}
