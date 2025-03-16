package views

import (
	toolsRepo "c2/repos/tools"
	"strings"

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
	originalData := map[string][]string{
		"": {},
	}
	categories := map[string]bool{}

	for _, tool := range tools {
		if _, ok := categories[tool.ToolCategoryName]; !ok {
			originalData[""] = append(originalData[""], tool.ToolCategoryName)
			originalData[tool.ToolCategoryName] = []string{tool.Name}
			categories[tool.ToolCategoryName] = true
			continue
		}
		originalData[tool.ToolCategoryName] = append(originalData[tool.ToolCategoryName], tool.Name)
	}

	// Filtered data. (starts as a copy of original)
	filteredData := make(map[string][]string)
	for k, v := range originalData {
		filteredData[k] = append([]string{}, v...) // Deep copy.
	}

	// Create the tree with callbacks.
	tree := widget.NewTree(
		// ChildUIDs: Returns children for a given node ID.
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			return filteredData[id]
		},
		// IsBranch: Determines if a node is a branch.
		func(id widget.TreeNodeID) bool {
			_, isBranch := filteredData[id]
			return isBranch
		},
		// CreateNode: Defines the appearance of nodes.
		func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("")
		},
		// UpdateNode: Sets the content of a node.
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(id)
		},
	)

	tree.OnSelected = func(name string) {
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

	// Search input.
	searchToolInput := widget.NewEntry()
	searchToolInput.SetPlaceHolder("Search Library")

	// Filter tree based on search input.
	searchToolInput.OnChanged = func(searchText string) {
		// Reset filteredData to original state.
		filteredData = make(map[string][]string)
		for k, v := range originalData {
			filteredData[k] = append([]string{}, v...)
		}

		if searchText == "" {
			// If search is empty, show all data.
			tree.Refresh()
			return
		}

		// Case-insensitive search.
		searchText = strings.ToLower(searchText)

		// Filter tools based on search text.
		matchedCategories := map[string]bool{}
		for category, toolsInCategory := range originalData {
			if category == "" {
				continue // Skip root
			}
			filteredTools := []string{}
			for _, toolName := range toolsInCategory {
				if strings.Contains(strings.ToLower(toolName), searchText) {
					filteredTools = append(filteredTools, toolName)
				}
			}
			if len(filteredTools) > 0 {
				filteredData[category] = filteredTools
				matchedCategories[category] = true
			} else {
				// Remove category if no matches.
				delete(filteredData, category)
			}
		}

		// Update root to only show categories with matches.
		filteredData[""] = []string{}
		for category := range matchedCategories {
			filteredData[""] = append(filteredData[""], category)
		}

		// Refresh the tree to reflect the filtered data.
		tree.CloseAllBranches() // Collapse all for clarity. (maybe not gonna keep it this way... shall see...)
		tree.Refresh()
	}

	return container.NewBorder(searchToolInput, nil, nil, nil, tree)
}
