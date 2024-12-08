package todayWeather

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/GopherGhaznix/FyneWeather/bundled"
	"github.com/GopherGhaznix/FyneWeather/config"
)

func View() fyne.CanvasObject {
	tempText := canvas.NewText("18°", theme.Color(theme.ColorNameForeground))
	tempText.FontSource = config.FontTekoBold
	tempText.TextSize = 65
	tempText.TextStyle = fyne.TextStyle{}

	DescriptionText := canvas.NewText("Thunderstom", theme.Color(theme.ColorNamePressed))
	DescriptionText.TextSize = 12

	return container.NewBorder(
		nil, nil,
		container.NewCenter(
			container.NewStack(
				tempText,
				container.NewBorder(nil, DescriptionText, nil, nil),
			),
		),
		container.NewGridWrap(fyne.NewSize(125, 90),
			canvas.NewImageFromResource(bundled.ResourceIsolatedThunderstormsSvg),
		),
	)
}
