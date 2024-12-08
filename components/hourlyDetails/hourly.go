package hourlyDetails

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
					timebasedcard.View("./assets/animated/isolated-thunderstorms.svg", "7 AM", "28°"),
					timebasedcard.View("./assets/animated/cloudy.svg", "7 AM", "28°"),
					timebasedcard.View("./assets/animated/cloudy-3-day.svg", "7 AM", "28°"),
					timebasedcard.View("./assets/animated/rain-and-sleet-mix.svg", "7 AM", "28°"),
					timebasedcard.View("./assets/animated/snowy-2.svg", "7 AM", "28°"),
				),
				_space, // just for decent amount of spacing
			),
		),
	)
}
