package components

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// LargeTextViewer is a custom Fyne widget for displaying large amounts of text efficiently using a List.
type LargeTextViewer struct {
	widget.BaseWidget
	lines []string     // The text content split into lines
	list  *widget.List // List widget to display the lines
}

// NewLargeTextViewer creates a new LargeTextViewer widget with the given text.
func NewLargeTextViewer(text string) *LargeTextViewer {
	viewer := &LargeTextViewer{
		lines: strings.Split(text, "\n"),
	}
	viewer.ExtendBaseWidget(viewer)

	// Create the List widget
	viewer.list = widget.NewList(
		// Length function: returns the total number of lines
		func() int {
			return len(viewer.lines)
		},
		// Create function: creates a template for each item
		func() fyne.CanvasObject {
			text := canvas.NewText("", theme.Color(theme.ColorNameForeground))
			text.TextSize = theme.TextSize()
			return text
		},
		// Update function: updates the item at the given index
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < 0 || id >= len(viewer.lines) {
				return // Prevent out-of-bounds access
			}
			text := obj.(*canvas.Text)
			text.Text = viewer.lines[id]
			text.Refresh()
		},
	)

	return viewer
}

// SetText updates the text content of the viewer.
func (v *LargeTextViewer) SetText(text string) {
	v.lines = strings.Split(text, "\n")
	v.list.Refresh()
}

// CreateRenderer creates the renderer for the LargeTextViewer.
func (v *LargeTextViewer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(v.list)
}

// Resize updates the widget's size.
func (v *LargeTextViewer) Resize(size fyne.Size) {
	v.BaseWidget.Resize(size)
	v.list.Resize(size)
}
