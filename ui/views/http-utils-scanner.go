package views

import (
	"slices"
	"ui/core"
	"ui/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	c2api "github.com/zarkones/xena-client"
)

func NewHttpScannerDialog(reqID int64) {
	w := core.App.NewWindow("Send Request For Scanning")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	nonSelectedAgents := make([]string, len(state.Agents))
	selectedAgents := []string{}

	for i := 0; i < len(state.Agents); i++ {
		nonSelectedAgents[i] = state.Agents[i].Hostname + " " + state.Agents[i].ID
	}

	var nonSelectedAgentsTable *widget.List
	var selectedAgentsTable *widget.List

	nonSelectedAgentsTable = widget.NewList(
		func() int {
			return len(nonSelectedAgents)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			agent := nonSelectedAgents[i]
			o.(*widget.Button).SetText(agent)
			o.(*widget.Button).OnTapped = func() {
				selectedAgents = append(selectedAgents, agent)
				nonSelectedAgents = slices.Delete(nonSelectedAgents, i, i+1)
				go nonSelectedAgentsTable.Refresh()
				go selectedAgentsTable.Refresh()
			}
		},
	)

	selectedAgentsTable = widget.NewList(
		func() int {
			return len(selectedAgents)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			agent := selectedAgents[i]
			o.(*widget.Button).SetText(agent)
			o.(*widget.Button).OnTapped = func() {
				nonSelectedAgents = append(nonSelectedAgents, agent)
				selectedAgents = slices.Delete(selectedAgents, i, i+1)
				go nonSelectedAgentsTable.Refresh()
				go selectedAgentsTable.Refresh()
			}
		},
	)

	confirmBtn := widget.NewButton("CONFIRM", func() {
		if err := c2api.InsertHttpScan(reqID, selectedAgents); err != nil {
			Alert("Failed to insert scan: " + err.Error())
			return
		}

		w.Close()
	})

	w.SetContent(container.NewBorder(
		// Top.
		container.NewVBox(
			widget.NewLabel("Select agents to perform scanning"),
			container.NewHBox(
				widget.NewLabel("Not selected:"),
				layout.NewSpacer(),
				widget.NewLabel("Selected:"),
			),
		),

		// Bottom.
		confirmBtn,

		// Left.
		nil,

		// Right.
		nil,

		// Center.
		container.NewHSplit(
			nonSelectedAgentsTable,
			selectedAgentsTable,
		),
	))

	w.Show()
}

func HttpUtilsScansTable() fyne.CanvasObject {
	// TODO: Open a dialog to ask which agent to use for scanning.
	return container.NewStack()
}
