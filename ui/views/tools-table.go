package views

import (
	toolsRepo "c2/repos/tools"

	dia "fyne.io/x/fyne/widget/diagramwidget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	xenaC2 "github.com/zarkones/xena-client"
)

var toolNames []string
var tools map[string]xenaC2.Tool = func() map[string]xenaC2.Tool {
	tools, err := toolsRepo.GetMultiple()
	if err != nil {
		// TODO: Uncomment: Warn("failed to load tools:" + err.Error())
		return nil
	}
	toolMap := make(map[string]xenaC2.Tool, len(tools))
	toolNames = make([]string, len(tools))
	for i, tool := range tools {
		toolMap[tool.Name] = tool
		toolNames[i] = tool.Name
	}
	return toolMap
}()

func NewToolsTableForPipeline(diagram *dia.DiagramWidget) (table fyne.CanvasObject) {
	t := map[string][]string{
		"": {},
	}

	categories := map[string]bool{}

	for _, tool := range tools {
		if _, ok := categories[tool.ToolCategoryName]; !ok {
			t[""] = append(t[""], tool.ToolCategoryName)
			t[tool.ToolCategoryName] = []string{tool.Name}
			categories[tool.ToolCategoryName] = true
			continue
		}
		t[tool.ToolCategoryName] = append(t[tool.ToolCategoryName], tool.Name)
	}

	StateTree := widget.NewTreeWithStrings(t)

	StateTree.OnSelected = func(name string) {
		go func() {
			if _, ok := categories[name]; ok {
				return
			}

			newNodeID := uuid.NewString()
			pos := fyne.NewPos(200, diagram.Size().Height/2)
			tool := tools[name]
			copiedInputs := make(map[string]xenaC2.ToolInput, len(tool.Inputs))
			for key, val := range tool.Inputs {
				copiedInputs[key] = val
			}
			currPipeSettings.Steps[newNodeID] = xenaC2.PipelineStep{
				ID:       newNodeID,
				Position: pos,
				Name:     name,
				Tool: xenaC2.Tool{
					ID:               tool.ID,
					Name:             tool.Name,
					Description:      tool.Description,
					ToolCategoryName: tool.ToolCategoryName,
					Inputs:           copiedInputs,
				},
			}

			ref := currPipeSettings.Steps[newNodeID]

			setStep(&ref)
		}()
	}

	searchToolInput := widget.NewEntry()
	searchToolInput.SetPlaceHolder("Search Library")

	return container.NewBorder(searchToolInput, nil, nil, nil, StateTree)
}
