package toolsRepo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	xenaC2 "github.com/zarkones/xena-client"
)

const TOOLS_DIR = "xena-tools"

func GetMultiple() (tools []xenaC2.Tool, err error) {
	toolFiles, err := os.ReadDir(TOOLS_DIR)
	if err != nil {
		return nil, err
	}

	for _, toolFile := range toolFiles {
		toolsJson, err := os.ReadFile(filepath.Join(TOOLS_DIR, toolFile.Name()))
		if err != nil {
			fmt.Println("Failed to read tools file:", toolFile, ", exception:", err)
			continue
		}
		var toolsBatch []xenaC2.Tool
		if err := json.Unmarshal(toolsJson, &toolsBatch); err != nil {
			if err != nil {
				fmt.Println("Failed to unmarshal tools file:", toolFile, ", exception:", err)
				continue
			}
		}
		tools = append(tools, toolsBatch...)
	}

	return tools, nil
}

func Insert(tool *xenaC2.Tool) (err error) {
	toolJson, err := json.Marshal(&tool)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(TOOLS_DIR, tool.Name), toolJson, 0777)
}
