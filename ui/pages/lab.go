package pages

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Lab() fyne.CanvasObject {
	xenaClientURL, _ := url.Parse("https://github.com/zarkones/xena-client")
	xenaAgentURL, _ := url.Parse("https://github.com/zarkones/xena-agent")
	xenaLinkclient := widget.NewHyperlink("github.com/zarkones/xena-client", xenaClientURL)
	xenaLinkagent := widget.NewHyperlink("github.com/zarkones/xena-agent", xenaAgentURL)
	return container.NewScroll(
		container.NewVBox(
			widget.NewLabel(`
BUILDING CUSTOM AGENTS

Integration into XENA ecosystem is simple.`),

			widget.NewLabel(`XENA C2 API HTTP Client Library:`),
			xenaLinkclient,

			widget.NewLabel(`Example agent made using the client library:`),
			xenaLinkagent,
		))
}
