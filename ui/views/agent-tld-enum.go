package views

import (
	"os"
	"strings"
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func AgentTldEnum(callback func(host string, wordlist []string)) {
	w := core.App.NewWindow("XENA: Top-Level Domain Enumeration")
	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()

	targetHostInput := widget.NewEntry()
	targetHostInput.SetText("example.com")

	wordlist := []string{}

	selectWordlistBtn := widget.NewButton("Select Wordlist", func() {
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

			wordlist = strings.Split(string(content), "\n")
		}, w)
		fd.Show()
	})

	runBtn := widget.NewButton("Run", func() {
		callback(targetHostInput.Text, wordlist)
		w.Close()
	})

	w.SetContent(container.NewVBox(
		targetHostInput,
		widget.NewLabel("If you don't select a wordlist the default one will be used."),
		selectWordlistBtn,
		runBtn,
	))

	w.Show()
}
