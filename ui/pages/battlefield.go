package pages

import (
	"bytes"
	"image/png"
	"math"
	"strings"
	"time"
	"ui/state"
	"ui/static"
	"ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	dia "fyne.io/x/fyne/widget/diagramwidget"
	xenaC2 "github.com/zarkones/xena-client"
	"golang.org/x/exp/slices"
)

var diagramWidget = dia.NewDiagramWidget("ENGAGEMENT")

var avatar *canvas.Image
var agentAvatar *canvas.Image
var osLogoWindows *canvas.Image
var osLogoLinux *canvas.Image
var osLogoDarwin *canvas.Image

const POS_OFFSET = 300

func generatePositions(initial fyne.Position, numPoints int) []fyne.Position {
	positions := make([]fyne.Position, numPoints)

	angle := 360.0 / float64(numPoints)

	for i := 0; i < numPoints; i++ {
		radians := angle * float64(i) * (math.Pi / 180.0)

		x := initial.X + float32(math.Cos(radians)*POS_OFFSET)
		y := initial.Y + float32(math.Sin(radians)*POS_OFFSET)

		positions[i] = fyne.Position{X: x, Y: y}
	}

	return positions
}

type TargetView struct {
	fyne.Container
	Node *dia.DiagramNode
}

func (mv *TargetView) Tapped(e *fyne.PointEvent) {
	agents, _ := xenaC2.GetAgents()
	menuItems := make([]*fyne.MenuItem, len(agents)+1)
	menuItems[0] = fyne.NewMenuItem("Inspect", func() {
		// TODO: Select the target in the inspector.
	})
	for i, agent := range agents {
		menuItems[i+1] = fyne.NewMenuItem("Atk with: "+agent.Hostname, func() {
			func(agentID string) {
				targetNode := *mv.Node

				targetID := strings.TrimPrefix(targetNode.GetDiagramElementID(), "TARGET:")
				if err := xenaC2.AttackTarget(agentID, targetID); err != nil {
					views.Warn("Failed To Issue An Attack Order To The C2 Server, exception: " + err.Error())
					return
				}
			}(agent.ID)
		})
	}

	m := fyne.NewMenu(
		"Attack Menu",
		menuItems...,
	)
	widget.ShowPopUpMenuAtPosition(
		m,
		fyne.CurrentApp().Driver().CanvasForObject(diagramWidget),
		e.AbsolutePosition,
	)
}

func newAgentNode(agent xenaC2.Agent) *fyne.Container {
	var osLogo *canvas.Image
	switch agent.OS {
	case "windows":
		osLogo = osLogoWindows
	case "linux":
		osLogo = osLogoLinux
	case "darwin":
		osLogo = osLogoDarwin
	default:
		osLogo = avatar
	}

	return container.NewVBox(
		// 	// Primary.
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewLabel(agent.Hostname),
			layout.NewSpacer(),
		),
		// ),

		agentAvatar,

		// BOTTOM.
		container.NewHBox(
			osLogo,

			widget.NewLabel(agent.IpAddress),

			widget.NewButtonWithIcon("", theme.MoreVerticalIcon(), func() {
				views.AgentWindow(agent)
			}),
		),
	)
}

func Engagement() fyne.CanvasObject {
	displayedAgentIDs := []string{}

	// C2 sprite.
	c2Img, _ := png.Decode(bytes.NewReader(static.C2))
	c2Sprite := canvas.NewImageFromImage(c2Img)
	c2Sprite.FillMode = canvas.ImageFillOriginal
	// XENA Logo on agents.
	xenalogoImg, _ := png.Decode(bytes.NewReader(static.XenaAvatar))
	avatar = canvas.NewImageFromImage(xenalogoImg)
	avatar.FillMode = canvas.ImageFillOriginal
	// Agent avatar.
	agentlogoImg, _ := png.Decode(bytes.NewReader(static.AgentAvatar))
	agentAvatar = canvas.NewImageFromImage(agentlogoImg)
	agentAvatar.FillMode = canvas.ImageFillOriginal
	// OS Avatars.
	osLogoWindowsImg, _ := png.Decode(bytes.NewReader(static.OsAvatarWindows))
	osLogoWindows = canvas.NewImageFromImage(osLogoWindowsImg)
	osLogoWindows.FillMode = canvas.ImageFillOriginal
	osLogoLinuxImg, _ := png.Decode(bytes.NewReader(static.OsAvatarLinux))
	osLogoLinux = canvas.NewImageFromImage(osLogoLinuxImg)
	osLogoLinux.FillMode = canvas.ImageFillOriginal
	osLogoDarwinImg, _ := png.Decode(bytes.NewReader(static.OsAvatarDarwin))
	osLogoDarwin = canvas.NewImageFromImage(osLogoDarwinImg)
	osLogoDarwin.FillMode = canvas.ImageFillOriginal

	// Main Diagram thingies.
	scrollContainer := container.NewScroll(diagramWidget)
	scrollContainer.Offset = fyne.NewPos(500, 500)

	// C2 Node.
	c2NodeID := "Command & Control Server"
	c2Node := dia.NewDiagramNode(
		diagramWidget,
		container.NewVBox(
			c2Sprite,
			widget.NewLabel(c2NodeID),
		),
		c2NodeID,
	)
	c2NodeX := float32(300)
	c2NodeY := float32(300)
	c2Node.Move(fyne.NewPos(c2NodeX, c2NodeY))
	c2Node.SetProperties(dia.DiagramElementProperties{
		StrokeWidth: 0,
	})
	c2Node.Refresh()

	updateDiagram := func() {
		state.Agents, _ = xenaC2.GetAgents()

		points := generatePositions(c2Node.Position(), len(state.Agents))

		for i, agent := range state.Agents {
			if slices.Contains(displayedAgentIDs, agent.ID) {
				continue
			}
			displayedAgentIDs = append(displayedAgentIDs, agent.ID)

			agentNodeID := "AGENT:" + agent.ID
			agentNode := dia.NewDiagramNode(diagramWidget, nil, agentNodeID)
			agentNode.SetProperties(dia.DiagramElementProperties{
				StrokeWidth: 0,
			})
			agentNode.Move(points[i])

			agentNode.SetInnerObject(newAgentNode(agent))

			agentLinkID := "NODE_LINK:" + c2NodeID + "->" + agentNodeID
			agentLink := dia.NewDiagramLink(diagramWidget, agentLinkID)
			agentLink.SetSourcePad(c2Node.GetEdgePad())
			agentLink.SetTargetPad(agentNode.GetEdgePad())
			agentLink.AddSourceDecoration(dia.NewArrowhead())
		}
	}

	go func() {
		go updateDiagram()
		for range time.Tick(time.Second * 5) {
			if state.AuthToken == "" {
				continue
			}
			updateDiagram()
		}
	}()

	return container.NewBorder(
		views.AssetsToolbar(),
		nil,
		nil,
		nil,
		scrollContainer,
	)
}
