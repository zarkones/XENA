package layouts

import (
	"ui/effects"
	"ui/pages"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Default() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Engagement", pages.Engagement()),
		container.NewTabItem("Pipelines", pages.Pipelines()),
		container.NewTabItem("Files", pages.Files()),
		container.NewTabItem("Http Utils", pages.HttpUtils()),
		container.NewTabItem("Shop", pages.Shop()),
		container.NewTabItem("Lab", pages.Lab()),
		container.NewTabItem("Settings", pages.Settings()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	return container.NewStack(
		effects.Gradient(tabs, false, false),
	)
}
