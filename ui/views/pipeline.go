package views

import (
	"bytes"
	"encoding/json"
	"image/png"
	"slices"
	"strings"
	"time"
	"ui/core"
	"ui/effects"
	"ui/state"
	"ui/static"

	"github.com/google/uuid"
	xenaC2 "github.com/zarkones/xena-client"
	"golang.design/x/clipboard"

	dia "fyne.io/x/fyne/widget/diagramwidget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const PIPE_LINK_POINTER = "->"
const PIPE_LINK_PREFIX = "NODE_LINK:"

func PipelineWindow(pipe xenaC2.Pipeline) {
	w := core.App.NewWindow("XENA Pipeline: " + pipe.Name)
	w.Resize(fyne.NewSize(AGENT_WINDOW_WIDTH, AGENT_WINDOW_HEIGHT))
	w.CenterOnScreen()

	w.SetContent(pipeline(pipe, &w))

	w.SetOnClosed(func() {
		inspector = container.NewVBox()
	})

	w.Show()
}

var diagramWidget *dia.DiagramWidget
var inspector = container.NewVBox()
var currPipeSettings xenaC2.PipelineSettings
var currPipeWindow *fyne.Window
var inspectedNodeID string

var pentagonSprite *canvas.Image = func() *canvas.Image {
	pentagonImg, _ := png.Decode(bytes.NewReader(static.Pentagon))
	pentagon := canvas.NewImageFromImage(pentagonImg)
	pentagon.FillMode = canvas.ImageFillOriginal
	return pentagon
}()

type NodeView struct {
	fyne.Container
	Node *dia.DiagramNode
}

func (mv *NodeView) Tapped(e *fyne.PointEvent) {
	node := *mv.Node
	nodeID := node.GetDiagramElementID()
	InspectToolPipeline(nodeID, inspector)
}

func (mv *NodeView) TappedSecondary(e *fyne.PointEvent) {
	node := *mv.Node
	nodeID := node.GetDiagramElementID()
	menuItems := []*fyne.MenuItem{
		fyne.NewMenuItem("Inspect", func() {
			InspectToolPipeline(nodeID, inspector)
		}),
	}
	if len(potentialLinkIDs) == 1 && nodeID == potentialLinkIDs[0] {
		menuItems = append(menuItems, fyne.NewMenuItem("Unselect", func() {
			potentialLinkIDs = []string{}
		}))
	} else {
		linkLabel := "Select"
		if len(potentialLinkIDs) != 0 {
			linkLabel = "Link Here"
		}
		menuItems = append(menuItems, fyne.NewMenuItem(linkLabel, func() {
			linkNodes(nodeID)
		}))
	}
	menuItems = append(menuItems, fyne.NewMenuItem("Delete", func() {
		delStep(nodeID)
	}))
	m := fyne.NewMenu(
		"Pipeline Node Menu",
		menuItems...,
	)
	widget.ShowPopUpMenuAtPosition(
		m,
		fyne.CurrentApp().Driver().CanvasForObject(diagramWidget),
		e.AbsolutePosition,
	)
}

var potentialLinkIDs = []string{}

func setLink(step xenaC2.PipelineStep) {
	for _, linkedTo := range step.LinkedTo {
		newLinkID := PIPE_LINK_PREFIX + step.ID + PIPE_LINK_POINTER + linkedTo
		newLink := dia.NewDiagramLink(diagramWidget, newLinkID)
		newLink.SetTargetPad(diagramWidget.GetDiagramNode(step.ID).GetEdgePad())
		newLink.SetSourcePad(diagramWidget.GetDiagramNode(linkedTo).GetEdgePad())
		newLink.AddSourceDecoration(dia.NewArrowhead())
	}
}

func setStep(step xenaC2.PipelineStep) {
	targetView := NodeView{}

	newToolNode := dia.NewDiagramNode(
		diagramWidget,
		container.NewVBox(
			container.NewHBox(
				layout.NewSpacer(),
				widget.NewLabel(step.Name),
				layout.NewSpacer(),
			),
			container.NewHBox(
				layout.NewSpacer(),
				container.NewStack(
					pentagonSprite,
					&targetView,
				),
				layout.NewSpacer(),
			),
		),
		step.ID,
	)
	targetView.Node = &newToolNode
	newToolNode.Move(step.Position)
	newToolNode.SetProperties(dia.DiagramElementProperties{
		StrokeWidth: 0,
	})
	newToolNode.Refresh()
}

func delStep(stepID string) {
	w := core.App.NewWindow("XENA Deletion Dialog")
	w.Resize(fyne.NewSize(300, 0))
	w.CenterOnScreen()

	w.SetContent(container.NewVBox(
		widget.NewLabel("Are you sure you want to delete the node?"),
		container.NewHBox(
			widget.NewButton("CANCEL", func() {
				w.Close()
			}),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("DELETE", theme.DeleteIcon(), func() {
				delete(currPipeSettings.Steps, stepID)
				if slices.Contains(potentialLinkIDs, stepID) {
					potentialLinkIDs = []string{}
				}
				if inspectedNodeID == stepID {
					inspector.RemoveAll()
					inspector.Refresh()
				}
				diagramWidget.RemoveElement(stepID)
				w.Close()
			}),
		),
	))

	w.Show()
}

func linkNodes(nodeID string) {
	switch len(potentialLinkIDs) {
	case 0:
		potentialLinkIDs = append(potentialLinkIDs, nodeID)
		return
	case 1:
		potentialLinkIDs = append(potentialLinkIDs, nodeID)
	}
	newLinkID := PIPE_LINK_PREFIX + potentialLinkIDs[0] + PIPE_LINK_POINTER + potentialLinkIDs[1]
	newLink := dia.NewDiagramLink(diagramWidget, newLinkID)
	newLink.SetTargetPad(diagramWidget.GetDiagramNode(potentialLinkIDs[0]).GetEdgePad())
	newLink.SetSourcePad(diagramWidget.GetDiagramNode(potentialLinkIDs[1]).GetEdgePad())
	newLink.AddSourceDecoration(dia.NewArrowhead())
	potentialLinkIDs = []string{}
}

func pipeline(pipeline xenaC2.Pipeline, w *fyne.Window) fyne.CanvasObject {
	diagramWidget = dia.NewDiagramWidget("PIPELINE:" + pipeline.ID)
	toolsTable := NewToolsTableForPipeline(diagramWidget)
	currPipeWindow = w

	// Parse the pipeline's settings.
	currPipeSettings = xenaC2.PipelineSettings{}
	if err := json.Unmarshal([]byte(pipeline.Settings), &currPipeSettings); err != nil {
	}
	if currPipeSettings.Input == nil {
		currPipeSettings.Input = map[string]string{}
	}
	if currPipeSettings.Steps == nil {
		currPipeSettings.Steps = map[string]xenaC2.PipelineStep{}
	}

	savePipeline := func() {
		nodes := diagramWidget.GetDiagramNodes()
		for _, node := range nodes {
			step, ok := currPipeSettings.Steps[node.GetDiagramElementID()]
			if !ok {
				continue
			}
			step.Position = node.Position()
			currPipeSettings.Steps[node.GetDiagramElementID()] = step
		}

		links := diagramWidget.GetDiagramLinks()
		for _, link := range links {
			ids := strings.Split(
				strings.TrimPrefix(link.GetDiagramElementID(), PIPE_LINK_PREFIX),
				PIPE_LINK_POINTER,
			)
			step, ok := currPipeSettings.Steps[ids[0]]
			if !ok {
				continue
			}
			if step.LinkedTo == nil {
				step.LinkedTo = []string{ids[1]}
			} else {
				step.LinkedTo = append(step.LinkedTo, ids[1])
			}
			currPipeSettings.Steps[ids[0]] = step
		}

		serializedSettings, err := json.Marshal(&currPipeSettings)
		if err != nil {
			Warn("Failed to Serialize the Pipeline, exception:" + err.Error())
			return
		}

		pipeline.Settings = string(serializedSettings)

		if err := xenaC2.UpsertPipeline(pipeline); err != nil {
			Warn("Failed to Save the Pipeline, exception:" + err.Error())
			return
		}

		state.Pipelines, _ = xenaC2.GetPipelines()
	}

	// BTN: RUN PIPELINE
	runPipelineBtn := widget.NewButtonWithIcon("RUN", theme.MailSendIcon(), func() {
		savePipeline()

		w := core.App.NewWindow("XENA Pipeline Execution: " + pipeline.Name)
		w.Resize(fyne.NewSize(420, 0))
		w.CenterOnScreen()

		selectedAgent := ""
		agentIDs := make([]string, len(state.Agents))
		for i := 0; i < len(agentIDs); i++ {
			agentIDs[i] = state.Agents[i].Hostname + " " + state.Agents[i].ID
		}

		agentSelect := widget.NewSelect(agentIDs, func(s string) {
			selectedAgent = strings.Split(s, " ")[1]
		})

		confirmBtn := widget.NewButton("CONFIRM", func() {
			if selectedAgent == "" {
				return
			}
			if err := xenaC2.ExecutePipeline(pipeline.ID, []string{selectedAgent}); err != nil {
				Warn("Error while executing the pipeline, exception: " + err.Error())
				return
			}

			w.Close()
		})

		singleAgentExecution := container.NewVBox(
			widget.NewLabel("Select an agent which will execute this pipeline."),
			agentSelect,
			confirmBtn,
		)

		executionTabs := container.NewAppTabs(
			container.NewTabItem("Single Agent Execution", singleAgentExecution),
			container.NewTabItem("Multi-Agent Execution", container.NewCenter(widget.NewLabel("Soon..."))),
		)

		w.SetContent(executionTabs)

		w.Show()
	})

	// BTN: RUN PIPELINE
	// schedulePipelineBtn := widget.NewButtonWithIcon("SCHEDULE", theme.HistoryIcon(), func() {})

	// UNDO & REDO BUTTONS
	// undoBtn := widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() {})
	// redoBtn := widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {})

	// BTN: SAVE PIPELINE
	savePipelineBtn := widget.NewButtonWithIcon("SAVE", theme.DocumentSaveIcon(), savePipeline)

	cInputs := container.NewVBox()
	newInput := func(inputLabel, inputValue string) {
		field := widget.NewEntry()
		field.OnChanged = func(s string) {
			currPipeSettings.Input[inputLabel] = s
		}
		field.SetText(inputValue)
		cInputs.Add(widget.NewLabel(inputLabel))
		cInputs.Add(field)
	}

	newInputLabel := widget.NewEntry()
	newInputLabel.SetPlaceHolder("Input Key")
	newInputValue := widget.NewEntry()
	newInputValue.SetPlaceHolder("Input Value")
	addInputBtn := widget.NewButton("ADD", func() {
		if newInputLabel.Text == "" {
			Notify("Validation Error", "Input Key cannot be empty.")
			return
		}
		if newInputValue.Text == "" {
			Notify("Validation Error", "Input Value cannot be empty.")
			return
		}
		currPipeSettings.Input[newInputLabel.Text] = newInputValue.Text
		newInput(newInputLabel.Text, newInputValue.Text)
		newInputLabel.Text = ""
		newInputValue.Text = ""
		newInputLabel.Refresh()
		newInputValue.Refresh()
	})

	cInputs.Add(container.NewBorder(
		// Top.
		nil,
		// Bottom.
		nil,
		// Left.
		nil,
		// Right.
		addInputBtn,
		// Primary.
		container.NewGridWithColumns(
			2,
			newInputLabel,
			newInputValue,
		),
	))
	cInputs.Add(widget.NewSeparator())

	for inputLabel, inputValue := range currPipeSettings.Input {
		newInput(inputLabel, inputValue)
	}

	for _, step := range currPipeSettings.Steps {
		setStep(step)
	}
	for _, step := range currPipeSettings.Steps {
		setLink(step)
	}

	libraryTabs := container.NewAppTabs(
		// Cyber security tools browsing and inspection.
		container.NewTabItem("Tools", effects.Gradient(toolsTable, true, false)),
		// Inputs/Variables of the pipeline.
		container.NewTabItem("Inputs", effects.Gradient(container.NewVScroll(cInputs), false, false)),
	)

	diagramScroll := container.NewScroll(diagramWidget)

	editorCont := container.NewBorder(
		// Top.
		effects.Gradient(
			container.NewHBox(
				widget.NewLabel("Name: "+pipeline.Name),
				widget.NewLabel("Description: "+pipeline.Desc),
				layout.NewSpacer(),
				// undoBtn,
				// redoBtn,
				layout.NewSpacer(),
				runPipelineBtn,
				// schedulePipelineBtn,
				savePipelineBtn,
			),
			true,
			true,
		),

		// Bottom.
		nil,

		// Left.
		nil,

		// Right.
		nil,

		// Primary - center.
		container.NewHSplit(
			libraryTabs,

			container.NewHSplit(
				diagramScroll,
				container.NewVBox(
					container.NewHBox(layout.NewSpacer(), widget.NewLabel("NODE INSPECTOR"), layout.NewSpacer()),
					widget.NewSeparator(),
					inspector,
				),
			),
		),
	)

	historyCont := pipelineRuns(pipeline.ID)

	mainTabs := container.NewAppTabs(
		container.NewTabItem("Editor", editorCont),
		container.NewTabItem("Runs", historyCont),
	)

	mainTabs.SetTabLocation(container.TabLocationLeading)

	return effects.Gradient(mainTabs, true, true)
}

func pipelineRuns(pipelineID string) *fyne.Container {
	var runDiagram *dia.DiagramWidget
	diagramCont := container.NewStack()
	console := container.NewStack()

	var pipelineRuns []xenaC2.PipelineRun
	var runsTable *widget.List

	updateRuns := func() {
		pipelineRuns, _ = xenaC2.GetPipelineRuns(pipelineID)
		runsTable.Refresh()
	}

	runsTable = widget.NewList(
		func() int {
			return len(pipelineRuns)
		},

		func() fyne.CanvasObject {
			return widget.NewButton("ffff-ffff-ffff-ffffffffffffffffffffff", nil)
		},

		func(i widget.ListItemID, o fyne.CanvasObject) {
			label := pipelineRuns[i].FinishedAtLabel
			if pipelineRuns[i].FinishedPipeline == "" {
				label += " RUNNING"
			} else {
				label += " FINISHED"
			}
			o.(*widget.Button).SetText(label)
			o.(*widget.Button).OnTapped = func() {
				runDiagram = dia.NewDiagramWidget("PIPELINE_RUN:" + pipelineID)

				defer func() {
					console.Refresh()
					runDiagram.Refresh()
					diagramCont.RemoveAll()
					diagramCont.Add(runDiagram)
					diagramCont.Refresh()
				}()

				// Clean up.
				console.RemoveAll()

				var pipeline xenaC2.Pipeline
				if err := json.Unmarshal([]byte(pipelineRuns[i].FinishedPipeline), &pipeline); err != nil {
					console.Add(widget.NewLabel("Error while deserializing the pipeline: " + err.Error()))
					return
				}

				var settings xenaC2.PipelineSettings
				if err := json.Unmarshal([]byte(pipeline.Settings), &settings); err != nil {
					console.Add(widget.NewLabel("Error while deserializing the pipeline's settings: " + err.Error()))
					return
				}

				for _, step := range settings.Steps {
					stepCopy := step

					newNode := dia.NewDiagramNode(runDiagram, container.NewVBox(
						widget.NewLabel(stepCopy.Tool.Name),
						pentagonSprite,
						widget.NewButton("INSPECT", func() {
							defer console.Refresh()
							console.RemoveAll()

							toolInputs := container.NewVBox()
							for name, input := range stepCopy.Tool.Inputs {
								toolInputs.Add(widget.NewRichTextFromMarkdown("# " + name + "\n### [" + input.Description + "]:\n" + input.Value))
								toolInputs.Add(widget.NewSeparator())
							}

							stdout := ""
							stderr := ""
							analysis := ""
							if stepCopy.Tool.Outputs != nil {
								stdout = stepCopy.Tool.Outputs["stdout"].Value
								stderr = stepCopy.Tool.Outputs["stderr"].Value
								analysis = stepCopy.Tool.Outputs["analysis"].Value
							}

							stdoutView := container.NewHScroll(widget.NewRichText(&widget.TextSegment{
								Style: widget.RichTextStyleCodeBlock,
								Text:  stdout,
							}))
							stderrView := container.NewHScroll(widget.NewRichText(&widget.TextSegment{
								Style: widget.RichTextStyleCodeBlock,
								Text:  stderr,
							}))
							analysisView := container.NewHScroll(widget.NewRichText(&widget.TextSegment{
								Style: widget.RichTextStyleCodeBlock,
								Text:  analysis,
							}))

							tabs := container.NewAppTabs(
								container.NewTabItem("STDOUT", container.NewVBox(
									container.NewHBox(
										widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
											clipboard.Write(clipboard.FmtText, []byte(stdout))
										}),
									),
									stdoutView,
								)),
								container.NewTabItem("STDERR", container.NewVBox(
									container.NewHBox(
										widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
											clipboard.Write(clipboard.FmtText, []byte(stderr))
										}),
									),
									stderrView,
								)),
								container.NewTabItem("ANALYSIS", container.NewVBox(
									container.NewHBox(
										widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
											clipboard.Write(clipboard.FmtText, []byte(analysis))
										}),
									),
									analysisView,
								)),
								container.NewTabItem("PROPERTIES", toolInputs),
							)
							console.Add(tabs)
						}),
					), stepCopy.ID)

					newNode.Move(stepCopy.Position)
					newNode.SetProperties(dia.DiagramElementProperties{
						StrokeWidth: 0,
					})
					newNode.Refresh()
				}

				for _, step := range settings.Steps {
					for _, linkedTo := range step.LinkedTo {
						newLink := dia.NewDiagramLink(runDiagram, uuid.NewString())
						newLink.SetTargetPad(runDiagram.GetDiagramNode(step.ID).GetEdgePad())
						newLink.SetSourcePad(runDiagram.GetDiagramNode(linkedTo).GetEdgePad())
						newLink.AddSourceDecoration(dia.NewArrowhead())
					}
				}
			}
		},
	)

	go func() {
		updateRuns()
		for range time.Tick(time.Second * 5) {
			updateRuns()
		}
	}()

	return container.NewBorder(
		nil,
		nil,
		nil,
		effects.Gradient(container.NewVScroll(runsTable), false, false),
		container.NewVSplit(
			container.NewScroll(diagramCont),
			effects.Gradient(container.NewVScroll(console), false, false),
		),
	)
}
