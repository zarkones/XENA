package views

import (
	"c2/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var requestsCh = make(chan int64, 9)

var editingRequests = map[int64]models.ProxyReq{}

func HttpUtilsEditor() fyne.CanvasObject {
	reqs := container.NewVBox()

	go func() {
		for reqID := range requestsCh {
			reqs.Add(widget.NewButton(strconv.Itoa(int(reqID)), func() {
				// TODO:
			}))
		}
	}()

	return container.NewBorder(
		nil,
		nil,
		reqs,
		nil,
	)
}
