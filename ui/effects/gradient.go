package effects

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var colorA = color.RGBA{74, 57, 97, 255}
var colorB = color.RGBA{40, 42, 54, 255} //BG

func Gradient[T fyne.CanvasObject](element T, horizontal, reverseColors bool) *fyne.Container {
	if reverseColors {
		tempColor := colorA
		colorA = colorB
		colorB = tempColor
	}

	gradient := func() fyne.CanvasObject {
		if horizontal {
			return canvas.NewHorizontalGradient(colorB, colorA)
		}
		return canvas.NewVerticalGradient(colorB, colorA)
	}()

	c := container.NewStack(gradient)
	c.Add(element)

	return c
}
