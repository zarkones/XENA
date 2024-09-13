package pages

import (
	"ui/effects"
	"ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func HttpUtils() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Logger", views.HttpUtilsLogger()),
		container.NewTabItem("Editor", views.HttpUtilsEditor()),
		// container.NewTabItem("Settings", views.HttpUtilsSettings()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	return container.NewStack(
		effects.Gradient(tabs, false, false),
	)
}
