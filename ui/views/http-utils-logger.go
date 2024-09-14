package views

import (
	"c2/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func HttpUtilsLogger() fyne.CanvasObject {
	traffic := [][]string{}

	t := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(traffic), 6
		},
		func() fyne.CanvasObject {
			r := widget.NewRichText(&widget.TextSegment{Text: ""})
			r.Truncation = fyne.TextTruncateEllipsis
			// actionBtn := widget.NewButton("test", func() {})
			// actionBtn.Hide()
			return container.NewStack(r) // , actionBtn)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			// if i.Col == 4 {
			// 	o.Resize(fyne.NewSize(40, o.MinSize().Height))
			// 	o.(*fyne.Container).Objects[1].Show()
			// 	o.(*fyne.Container).Objects[1].Refresh()
			// } else {
			// if i.Col < 4 {
			if o.(*fyne.Container).Objects[0].(*widget.RichText).Segments[0].Textual() != traffic[i.Row][i.Col] {
				o.(*fyne.Container).Objects[0].(*widget.RichText).Segments[0] = &widget.TextSegment{Text: traffic[i.Row][i.Col]}
				o.Refresh()
			}
			// o.(*fyne.Container).Objects[0].Refresh()
			// }
		})

	t.SetColumnWidth(0, 64)
	t.SetColumnWidth(1, 64)
	t.SetColumnWidth(2, 80)
	t.SetColumnWidth(3, 300)
	t.SetColumnWidth(4, 300)
	t.SetColumnWidth(5, 300)

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
				return "METHOD"
			case 2:
				return "LENGTH"
			case 3:
				return "HOST"
			case 4:
				return "PATH"
			case 5:
				return "TIME"
			case 6:
				return ""
			}
		}()
		o.(*widget.Label).SetText(txt)
	}

	page := 0
	lastPage := -1
	lastLen := -1

	pageLabel := widget.NewLabel("")

	updateTraffic := func(page int) {
		// TODO: Use XENA C2 HTTP Client instead.
		pageLabel.SetText(strconv.Itoa(page + 1))

		resp, err := http.DefaultClient.Get("http://127.0.0.1:8080/v1/proxy/traffic?page=" + strconv.Itoa(page))
		if err != nil {
			fmt.Println("http.DefaultClient.Get:", err)
			return
		}

		if resp.StatusCode == http.StatusNoContent {
			traffic = [][]string{}
			return
		}

		var reqs []models.ProxyReq

		if err := json.NewDecoder(resp.Body).Decode(&reqs); err != nil {
			fmt.Println("json.NewDecoder:", err)
			return
		}

		if lastLen == len(reqs) && lastPage == page {
			return
		}

		traffic = make([][]string, len(reqs))

		for i := 0; i < len(reqs); i++ {
			traffic[i] = []string{
				strconv.Itoa(int(reqs[i].ID)),
				reqs[i].Method,
				strconv.Itoa(reqs[i].Length),
				reqs[i].Host,
				reqs[i].Path,
				reqs[i].Time.String(),
			}
		}

		lastLen = len(reqs)
		lastPage = page
	}

	go updateTraffic(page)

	return container.NewBorder(
		// container.NewScroll(trafficCont),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
				if page == 0 {
					return
				}
				page--
				updateTraffic(page)
				t.Refresh()
			}),
			pageLabel,
			widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
				page++
				updateTraffic(page)
				t.Refresh()
			}),
		),
		nil, nil, nil,
		t,
	)
}
