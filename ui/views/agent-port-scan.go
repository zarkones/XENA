package views

import (
	"strconv"
	"strings"
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AgentPortScan(callback func(host string, ports []int)) {
	w := core.App.NewWindow("XENA: Port Scan")
	w.Resize(fyne.NewSize(300, 0))
	w.CenterOnScreen()

	targetHostInput := widget.NewEntry()
	targetHostInput.SetPlaceHolder("example.com")

	targetPortInput := widget.NewEntry()
	targetPortInput.SetPlaceHolder("Comma Separated Ports")

	runBtn := widget.NewButton("Run", func() {
		ports := []int{}
		for _, rawPort := range strings.Split(targetPortInput.Text, ",") {
			port, err := strconv.Atoi(targetPortInput.Text)
			if err != nil {
				Warn("port number malformed: " + rawPort + "    " + err.Error())
				continue
			}
			ports = append(ports, port)
		}

		callback(targetHostInput.Text, ports)
		w.Close()
	})

	w.SetContent(container.NewVBox(
		targetHostInput,
		targetPortInput,
		runBtn,
	))

	w.Show()
}
