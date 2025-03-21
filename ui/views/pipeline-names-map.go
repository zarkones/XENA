package views

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

// This map is used to update the diagram in real tiem as you update the friendly name in the inspector.

var namesMap = map[string]*widget.Label{}

func setStepName(id, friendlyName string) {
	if _, ok := namesMap[id]; !ok {
		fmt.Println("friendly name not found for step:", id)
		return
	}
	namesMap[id].SetText(friendlyName)
}
