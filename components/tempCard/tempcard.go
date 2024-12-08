package tempCard

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/GopherGhaznix/FyneWeather/config"
)

func View() fyne.CanvasObject {
	tempText := canvas.NewText("18°", theme.Color(theme.ColorNameForeground))
	tempText.FontSource = config.FontTekoBold
	tempText.TextSize = 70
	tempText.TextStyle = fyne.TextStyle{}

	DescriptionText := canvas.NewText("Thunderstom", theme.Color(theme.ColorNamePressed))
	DescriptionText.TextSize = 12

	return container.NewVBox(
		container.NewCenter(
			container.NewStack(
				tempText,
				container.NewBorder(nil, DescriptionText, nil, nil),
			),
		),
	)
}
