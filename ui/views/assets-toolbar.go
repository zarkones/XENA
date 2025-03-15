package views

import (
	"common/builder"
	"ui/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func AssetsToolbar() fyne.CanvasObject {
	toolbar := container.NewHBox(
		widget.NewButtonWithIcon("NEW AGENT", theme.ContentAddIcon(), func() {
			go func() {
				w := core.App.NewWindow("Build Agent")
				w.Resize(fyne.NewSize(core.WIN_WIDTH, core.WIN_HEIGHT))
				w.SetContent(container.NewVScroll(builder.BuildAgent(w, w.Close)))
				w.CenterOnScreen()
				w.Show()
			}()
		}),
	)

	return toolbar
}
