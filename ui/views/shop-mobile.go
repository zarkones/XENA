package views

import (
	"ui/data"
	"ui/effects"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShopXenaMobile() fyne.CanvasObject {
	return container.NewVScroll(
		effects.Gradient(
			container.NewVBox(
				container.NewHBox(
					layout.NewSpacer(),
					widget.NewRichText(
						&widget.TextSegment{
							Style: widget.RichTextStyleHeading,
							Text:  "Introducing XENA: Mobile",
						},
						&widget.TextSegment{
							Style: widget.RichTextStyleCodeBlock,
							Text:  "The most potent cyber-security toolkit running on Android.",
						},
					),
					layout.NewSpacer(),
				),

				container.NewHBox(
					layout.NewSpacer(),
					widget.NewHyperlink("Buy Now", data.XenaMobileURL),
					layout.NewSpacer(),
				),

				container.NewVBox(
					container.NewHBox(
						layout.NewSpacer(),
						widget.NewRichText(
							&widget.TextSegment{
								Style: widget.RichTextStyleHeading,
								Text:  "FEATURES",
							},
							&widget.TextSegment{
								Style: widget.RichTextStyleSubHeading,
								Text:  "Port Scanner: Discover open ports on hosts.",
							},
							&widget.TextSegment{
								Style: widget.RichTextStyleSubHeading,
								Text:  "Subdomain Enumeration: Lower the visibility gap with performant multi-threaded asset discovery.",
							},
							&widget.TextSegment{
								Style: widget.RichTextStyleSubHeading,
								Text:  "HTTP Editor: Craft and send HTTP web requests with complete control.",
							},
							&widget.TextSegment{
								Style: widget.RichTextStyleSubHeading,
								Text:  "TLD Enumeration: Discover websites using efficient top-level domain enumeration.",
							},
							&widget.TextSegment{
								Style: widget.RichTextStyleSubHeading,
								Text:  "Web Crawler: Find links in a webpage.",
							},
							&widget.TextSegment{
								Style: widget.RichTextStyleSubHeading,
								Text:  "Reverse DNS: Locate websites residing on a specific IP address.",
							},
						),
						layout.NewSpacer(),
					),
				),

				container.NewHBox(
					layout.NewSpacer(),
					widget.NewHyperlink("Buy Now", data.XenaMobileURL),
					layout.NewSpacer(),
				),
			),
			false, false),
	)
}
