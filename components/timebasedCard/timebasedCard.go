package timebasedcard

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/GopherGhaznix/FyneWeather/config"
)

func View(image *fyne.StaticResource, hour, temp string) fyne.CanvasObject {

	hourText := canvas.NewText(hour, theme.Color(theme.ColorNamePressed))
	hourText.TextSize = 12

	tempText := canvas.NewText(temp, theme.Color(theme.ColorNameForeground))
	tempText.FontSource = config.FontTekoNormal
	tempText.TextSize = 18

	BG := canvas.NewRectangle(color.RGBA{R: 32, G: 35, B: 41, A: 255})
	BG.CornerRadius = 15

	return container.NewStack(

		BG,

		container.NewCenter(
			container.NewPadded(
				container.NewBorder(
					container.NewCenter(hourText),
					container.NewCenter(tempText),
					canvas.NewText("-", color.Transparent), // just for decent amount of spacing
					canvas.NewText("-", color.Transparent), // just for decent amount of spacing
					container.NewCenter(container.NewGridWrap(
						fyne.NewSize(55, 42),
						canvas.NewImageFromResource(image),
					)),
				),
			),
		),
	)
}
