package views

import (
	"c2/models"
	"common/netstack"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var editorReqCh = make(chan int64, 9)

func NewProxiedReqView(req models.ProxyReq) fyne.CanvasObject {
	reqEntry := widget.NewMultiLineEntry()
	reqEntry.SetText(req.RawReq)
	reqEntry.Wrapping = fyne.TextWrapWord

	respEntry := widget.NewRichText(&widget.TextSegment{Text: req.RawResp})
	// respEntry.Wrapping = fyne.TextWrapBreak

	secure := false
	secureConnCheck := widget.NewCheck("Secure", func(b bool) {
		secure = b
	})

	allowInsecure := false
	allowInsecureCheck := widget.NewCheck("Skip Cert. Check", func(b bool) {
		allowInsecure = b
	})

	timeoutEntry := widget.NewEntry()
	timeoutEntry.SetText("30")
	timeoutEntry.SetPlaceHolder("Timeout In Seconds")
	timeoutEntry.OnChanged = func(s string) {
		clean := ""
		for _, ch := range s {
			char := string(ch)
			if strings.Contains("0123456789", char) {
				clean += char
			}
		}
		timeoutEntry.SetText(clean)
	}

	hostEntry := widget.NewEntry()
	hostEntry.SetText(req.Host)
	if !strings.Contains(req.Host, ":") {
		hostEntry.SetText(net.JoinHostPort(req.Host, "443"))
	}

	sendReqBtn := widget.NewButton("SEND", func() {
		go func() {
			t, _ := strconv.Atoi(timeoutEntry.Text)
			timeout := time.Second * time.Duration(t)
			resp, err := netstack.Send(hostEntry.Text, reqEntry.Text, secure, allowInsecure, timeout)
			if err != nil {
				respEntry.Segments[0] = &widget.TextSegment{Text: err.Error()}
				return
			}

			respEntry.Segments[0] = &widget.TextSegment{Text: string(resp)}
			respEntry.Refresh()
		}()
	})

	topBar := container.NewGridWithColumns(4,
		hostEntry,
		container.NewHBox(sendReqBtn, secureConnCheck, allowInsecureCheck),
		container.NewGridWithColumns(4, timeoutEntry, layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer()),
		layout.NewSpacer(),
	)

	return container.NewBorder(
		topBar,
		nil,
		nil,
		nil,
		container.NewHSplit(container.NewScroll(reqEntry), container.NewScroll(respEntry)),
	)
}

func HttpUtilsEditor() fyne.CanvasObject {
	tabs := container.NewAppTabs()
	tabs.SetTabLocation(container.TabLocationTop)

	go func() {
		for reqID := range editorReqCh {
			id := fmt.Sprint(reqID)

			// TODO: Use XENA HTTP client.
			req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/v1/proxy/traffic/"+id, nil)
			if err != nil {
				fmt.Println("http editor: http.NewRequest:", err)
				continue
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println("http editor: http.DefaultClient.Do:", err)
				continue
			}

			var proxiedReq models.ProxyReq

			if err := json.NewDecoder(resp.Body).Decode(&proxiedReq); err != nil {
				fmt.Println("http editor: json.NewDecoder().Decode:", err)
				continue
			}

			tabs.Append(container.NewTabItem(id, NewProxiedReqView(proxiedReq)))
		}
	}()

	return tabs
}
