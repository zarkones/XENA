package views

import (
	"time"
	"ui/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func AgentsTable() fyne.CanvasObject {
	agents, err := xenaC2.GetAgents()
	if err != nil {
		Warn("failed to load agents:" + err.Error())
		return widget.NewLabel(err.Error())
	}

	table := widget.NewList(
		func() int {
			return len(agents)
		},

		func() fyne.CanvasObject {
			return widget.NewButton("template", nil)
		},

		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Button).SetText(agents[i].Hostname + " | " + agents[i].OS + " " + agents[i].Arch)
			o.(*widget.Button).OnTapped = func() {
				AgentWindow(agents[i])
			}
		},
	)

	go func() {
		for range time.Tick(time.Second * 5) {
			if state.AuthToken == "" {
				continue
			}
			agents, _ = xenaC2.GetAgents()
			table.Refresh()
		}
	}()

	return table
}
