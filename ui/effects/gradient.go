package effects

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Gradient[T fyne.CanvasObject](element T, horizontal, reverseColors bool) *fyne.Container {
	colorA := color.RGBA{74, 57, 97, 255}
	colorB := color.RGBA{40, 42, 54, 255} //BG

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
