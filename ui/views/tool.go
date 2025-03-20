package views

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	xenaC2 "github.com/zarkones/xena-client"
)

func InspectToolPipeline(nodeID string, inspector *fyne.Container) {
	defer inspector.Refresh()
	inspector.RemoveAll()
	inspectedNodeID = nodeID

	step, ok := currPipeSettings.Steps[nodeID]
	if !ok {
		inspector.Add(widget.NewLabel("Error: step " + nodeID + " not found"))
		return
	}

	friendlyNameInput := widget.NewEntry()
	friendlyNameInput.SetText(step.FriendlyName)
	friendlyNameInput.OnChanged = func(s string) {
		if len(s) == 0 {
			setStepName(step.ID, step.Name)
		} else {
			setStepName(step.ID, s)
		}

		currPipeSettings.Steps[nodeID] = xenaC2.PipelineStep{
			ID:           currPipeSettings.Steps[nodeID].ID,
			Name:         currPipeSettings.Steps[nodeID].Name,
			FriendlyName: s,
			Position:     currPipeSettings.Steps[nodeID].Position,
			Tool:         currPipeSettings.Steps[nodeID].Tool,
			LinkedTo:     currPipeSettings.Steps[nodeID].LinkedTo,
		}
	}

	inspector.Add(widget.NewLabel("Name: " + step.Tool.Name))
	inspector.Add(container.NewVBox(widget.NewLabel("Friendly Name:"), friendlyNameInput))
	inspector.Add(widget.NewLabel("Category: " + step.Tool.ToolCategoryName))
	inspector.Add(container.NewHScroll(widget.NewLabel("Description: " + step.Tool.Description)))

	inspector.Add(widget.NewSeparator())

	inspector.Add(container.NewHBox(layout.NewSpacer(), widget.NewLabel("INPUTS"), layout.NewSpacer()))
	inputs := step.Tool.Inputs
	for inputName, input := range inputs {
		// "OnChanged" callback would read the wrong value when called,
		// as the "inputName" would be set to last element of slice.
		inputName2 := inputName

		switch input.Type {

		case xenaC2.TOOL_IO_TYPE_STRING:
			newInput := widget.NewEntry()
			newInput.SetPlaceHolder(input.Description)
			newInput.SetText(currPipeSettings.Steps[nodeID].Tool.Inputs[inputName2].Value)
			newInput.OnChanged = func(s string) {
				currStepInput := currPipeSettings.Steps[nodeID].Tool.Inputs[inputName2]
				currStepInput.Value = s
				currPipeSettings.Steps[nodeID].Tool.Inputs[inputName2] = currStepInput
			}
			inspector.Add(widget.NewLabel(inputName2))
			inspector.Add(newInput)

		case xenaC2.TOOL_IO_TYPE_FILE:
			selectFileBtn := widget.NewButton("Select File", func() {
				fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
					if uc == nil {
						return
					}
					if err != nil {
						Alert("Failed to Open Dialog, exception: " + err.Error())
						return
					}
					sourcePath := uc.URI().Path()
					// sourceName := uc.URI().Name()
					content, err := os.ReadFile(sourcePath)
					if err != nil {
						Alert("Failed to Read File, exception: " + err.Error())
						return
					}

					currStepInput := currPipeSettings.Steps[nodeID].Tool.Inputs[inputName2]
					currStepInput.Value = string(content)
					currPipeSettings.Steps[nodeID].Tool.Inputs[inputName2] = currStepInput
				}, *currPipeWindow)
				fd.Show()
			})
			inspector.Add(widget.NewLabel(inputName2))
			inspector.Add(selectFileBtn)

		default:
			// We skip on unknown input type.
			// TODO: Alert or something.
			continue
		}
	}
}
