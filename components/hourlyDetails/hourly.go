package hourlyDetails

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/GopherGhaznix/FyneWeather/bundled"
	timebasedcard "github.com/GopherGhaznix/FyneWeather/components/timebasedCard"
)

func View() fyne.CanvasObject {

	_space := canvas.NewText("-", color.Transparent) // just for decent amount of spacing
	_space.TextSize = 5

	todayLabel := canvas.NewText("Today", theme.Color(theme.ColorNameForeground))

	return container.NewVBox(
		container.NewPadded(container.NewPadded(todayLabel)),
		container.NewHScroll(
			container.NewVBox(
				container.NewHBox(
					timebasedcard.View(bundled.ResourceIsolatedThunderstormsSvg, "7 AM", "28°"),
					timebasedcard.View(bundled.ResourceCloudySvg, "7 AM", "28°"),
					timebasedcard.View(bundled.ResourceCloudy3DaySvg, "7 AM", "28°"),
					timebasedcard.View(bundled.ResourceRainAndSleetMixSvg, "7 AM", "28°"),
					timebasedcard.View(bundled.ResourceSnowy2Svg, "7 AM", "28°"),
				),
				_space, // just for decent amount of spacing
			),
		),
	)
}
