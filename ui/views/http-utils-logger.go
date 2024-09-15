package views

import (
	"c2/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TableCell struct {
	widget.Label
	ReqID int64
}

func NewTableCell(txt string) *TableCell {
	l := &TableCell{}
	l.ExtendBaseWidget(l)
	return l
}

func (t *TableCell) Tapped(_ *fyne.PointEvent) {
}

func (t *TableCell) TappedSecondary(e *fyne.PointEvent) {
	sendToEditor := fyne.NewMenuItem("Send To Editor", func() {
		requestsCh <- t.ReqID
	})
	menu := fyne.NewMenu("Request Menu", sendToEditor)

	widget.ShowPopUpMenuAtPosition(
		menu,
		core.MainW.Canvas(),
		e.AbsolutePosition,
	)
}

func HttpUtilsLogger() fyne.CanvasObject {
	traffic := [][]string{}

	t := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(traffic), 9
		},
		func() fyne.CanvasObject {
			r := NewTableCell("")
			r.Truncation = fyne.TextTruncateEllipsis
			return r
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if o.(*TableCell).Text != traffic[i.Row][i.Col] {
				reqID, _ := strconv.Atoi(traffic[i.Row][0])
				o.(*TableCell).ReqID = int64(reqID)
				o.(*TableCell).SetText(traffic[i.Row][i.Col])
				o.(*TableCell).Refresh()
				o.Refresh()
			}
		})

	t.SetColumnWidth(0, 64)
	t.SetColumnWidth(1, 64)
	t.SetColumnWidth(2, 64)
	t.SetColumnWidth(3, 64)
	t.SetColumnWidth(4, 80)
	t.SetColumnWidth(5, 150)
	t.SetColumnWidth(6, 250)
	t.SetColumnWidth(7, 250)
	t.SetColumnWidth(8, 250)

	t.ShowHeaderColumn = false

	t.CreateHeader = func() fyne.CanvasObject {
		rTxt := widget.NewLabel("123123123")
		rTxt.Truncation = fyne.TextTruncateEllipsis
		return rTxt
	}
	t.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {
		txt := func() string {
			switch id.Col {
			default:
				return strconv.Itoa(id.Col)
			case 0:
				return "ID"
			case 1:
				return "STATUS"
			case 2:
				return "METHOD"
			case 3:
				return "REQ_LENGTH"
			case 4:
				return "RESP_LENGTH"
			case 5:
				return "HOST"
			case 6:
				return "PATH"
			case 7:
				return "QUERY"
			case 8:
				return "TIME"
			case 9:
				return ""
			}
		}()
		o.(*widget.Label).SetText(txt)
	}

	filterSelection := widget.NewSelect([]string{"STATUS", "METHOD", "REQ_LENGTH", "RESP_LENGTH", "HOST", "PATH", "QUERY", "TIME"}, func(s string) {})
	filterSelection.SetSelected("TIME")

	orderDirection := widget.NewSelect([]string{"ASCENDING", "DESCENDING"}, func(s string) {})
	orderDirection.SetSelected("ASCENDING")

	search := widget.NewEntry()
	search.SetPlaceHolder("Search...")
	search.Resize(fyne.NewSize(200, search.MinSize().Height))

	page := 0
	lastPage := -1
	lastLen := -1
	lastFilter := ""
	lastOrderDir := ""

	pageLabel := widget.NewLabel("")

	updateTraffic := func(page int) {
		if len(search.Text) != 0 {
			defer search.Enable()
			search.Disable()
		}
		// TODO: Use XENA C2 HTTP Client instead.
		pageLabel.SetText(strconv.Itoa(page + 1))

		resp, err := http.DefaultClient.Get("http://127.0.0.1:8080/v1/proxy/traffic?page=" + strconv.Itoa(page) + "&orderBy=" + filterSelection.Selected + "&order=" + orderDirection.Selected + "&search=" + search.Text)
		if err != nil {
			fmt.Println("http.DefaultClient.Get:", err)
			return
		}

		if resp.StatusCode == http.StatusNoContent {
			traffic = [][]string{}
			lastPage = -1
			lastLen = -1
			return
		}

		var reqs []models.ProxyReq

		if err := json.NewDecoder(resp.Body).Decode(&reqs); err != nil {
			fmt.Println("json.NewDecoder:", err)
			return
		}

		if lastLen == len(reqs) && lastPage == page && lastFilter == filterSelection.Selected && lastOrderDir == orderDirection.Selected {
			return
		}

		traffic = make([][]string, len(reqs))

		for i := 0; i < len(reqs); i++ {
			traffic[i] = []string{
				strconv.Itoa(int(reqs[i].ID)),
				strconv.Itoa(reqs[i].Status),
				reqs[i].Method,
				strconv.Itoa(reqs[i].ReqLength),
				strconv.Itoa(reqs[i].RespLength),
				reqs[i].Host,
				reqs[i].Path,
				reqs[i].Query,
				reqs[i].Time.String(),
			}
		}

		lastLen = len(reqs)
		lastPage = page
		lastFilter = filterSelection.Selected
		lastOrderDir = orderDirection.Selected
		t.Refresh()
	}

	filterSelection.OnChanged = func(s string) {
		updateTraffic(page)
	}
	orderDirection.OnChanged = func(s string) {
		updateTraffic(page)
	}

	search.OnSubmitted = func(s string) {
		updateTraffic(page)
	}

	go updateTraffic(page)

	return container.NewBorder(
		container.NewGridWithColumns(3,
			container.NewHBox(
				filterSelection,
				orderDirection,
				layout.NewSpacer(),
			),
			search,
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
					if page == 0 {
						return
					}
					page--
					updateTraffic(page)
				}),
				pageLabel,
				widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
					page++
					updateTraffic(page)
				}),
			),
		),
		nil, nil, nil,
		t,
	)
}
