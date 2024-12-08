package infobox

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func View(image *fyne.StaticResource, label, value string) fyne.CanvasObject {

	return container.NewPadded(
		container.NewPadded(
			container.NewPadded(
				container.NewVBox(
					container.NewCenter(container.NewGridWrap(
						fyne.NewSize(32, 32),
						canvas.NewImageFromResource(image),
					)),
					container.NewCenter(canvas.NewText(value, theme.Color(theme.ColorNameForeground))),
					container.NewCenter(canvas.NewText(label, theme.Color(theme.ColorNamePressed))),
				),
			),
		),
	)
}
