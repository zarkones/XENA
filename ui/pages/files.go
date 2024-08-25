package pages

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"ui/core"
	"ui/state"
	"ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func Files() fyne.CanvasObject {
	filesTable := widget.NewList(
		func() int {
			return len(state.Files)
		},

		func() fyne.CanvasObject {
			return widget.NewCard("asdasdasdsad", "adasdasdasdasdasdaasdasdasdasdasd", widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {}))
		},

		func(i widget.ListItemID, o fyne.CanvasObject) {
			file := state.Files[i]
			o.(*widget.Card).SetTitle(file.OriginalName)
			if file.Uploaded {
				o.(*widget.Card).SetSubTitle("Uploaded By Agent: " + file.UploadedByAgentID + ", Uploaded At: " + file.UploadedAt)
				o.(*widget.Card).SetContent(container.NewHBox(layout.NewSpacer(), widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {
					go downloadFile(file)
				})))
			} else {
				o.(*widget.Card).SetSubTitle("Requested Of Agent: " + file.UploadedByAgentID + ", Not Yet Uploaded")
				o.(*widget.Card).SetContent(nil)
			}
		},
	)

	updateFiles := func() {
		state.Files, _ = xenaC2.GetFiles()
		filesTable.Refresh()
	}

	go func() {
		updateFiles()
		for range time.Tick(time.Second * 10) {
			if state.AuthToken == "" {
				continue
			}
			updateFiles()
		}
	}()

	return filesTable
}

func downloadFile(file xenaC2.File) (err error) {
	wg := sync.WaitGroup{}

	_, orgFileName := filepath.Split(file.OriginalName)

	var fileContent []byte
	filePath := ""

	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		fileContent, err = xenaC2.DownloadFile(file.ID)
		if err != nil {
			views.Warn("File '" + file.OriginalName + "' failed to be saved, exception: " + err.Error())
		}
	}()

	wg.Add(1)
	fd := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		defer wg.Done()

		if err != nil || uc == nil {
			return
		}

		filePath = uc.URI().String()

		uc.Close()
	}, core.MainW)
	fd.SetFileName(orgFileName)
	fd.Show()

	wg.Wait()

	// Abort the operation if it was cancelled.
	if len(filePath) == 0 {
		return nil
	}
	if len(fileContent) == 0 {
		views.Warn("'" + orgFileName + "' is has no content")
		return nil
	}

	filePath = strings.TrimPrefix(filePath, "file://")

	if err := os.WriteFile(filePath, fileContent, 0777); err != nil {
		return err
	}

	return nil
}
