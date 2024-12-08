package lowhightempcard

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func View(label, image, low, high string) fyne.CanvasObject {

	lowTempText := canvas.NewText(low, theme.Color(theme.ColorNamePressed))
	lowTempText.TextSize = 15

	highTempText := canvas.NewText(high, theme.Color(theme.ColorNameForeground))
	highTempText.TextSize = 15

	return container.NewBorder(nil, nil,
		container.NewGridWrap(fyne.NewSize(85, 20), widget.NewLabelWithStyle(label, fyne.TextAlignLeading, fyne.TextStyle{})),
		container.NewCenter(container.NewGridWrap(
			fyne.NewSize(38, 28),
			canvas.NewImageFromFile(image),
		)),
		container.NewAdaptiveGrid(2,
			container.NewCenter(lowTempText),
			container.NewCenter(highTempText),
		),
	)
}
