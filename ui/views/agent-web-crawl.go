package views

import (
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AgentWebCrawl(callback func(domain string)) {
	w := core.App.NewWindow("XENA: Web Crawl")
	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()

	targetDomainInput := widget.NewEntry()
	targetDomainInput.SetText("example.com")

	runBtn := widget.NewButton("Run", func() {
		callback(targetDomainInput.Text)
		w.Close()
	})

	w.SetContent(container.NewVBox(
		targetDomainInput,
		runBtn,
	))

	w.Show()
}
