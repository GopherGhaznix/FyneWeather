package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/GopherGhaznix/FyneWeather/bundled"
	"github.com/GopherGhaznix/FyneWeather/components/headerscreen2"
	lowhightempcard "github.com/GopherGhaznix/FyneWeather/components/lowHighTempCard"
	"github.com/GopherGhaznix/FyneWeather/components/tempCard"
	"github.com/GopherGhaznix/FyneWeather/components/weatherInfo"
)

func WeeklyScreen() fyne.CanvasObject {

	return container.NewScroll(
		container.NewPadded(
			container.NewPadded(
				container.NewBorder(
					headerscreen2.View(),
					nil, nil, nil, container.NewPadded(
						container.NewVBox(
							tempCard.View(),
							canvas.NewText("", color.Transparent), // just for decent amount of spacing
							weatherInfo.View(),
							container.NewVBox(
								lowhightempcard.View("Today", "18°", "21°", bundled.ResourceIsolatedThunderstormsSvg),
								lowhightempcard.View("Thesday", "13°", "24°", bundled.ResourceIsolatedThunderstormsSvg),
								lowhightempcard.View("Wednesday", "23°", "31°", bundled.ResourceIsolatedThunderstormsSvg),
								lowhightempcard.View("Thusday", "15°", "29°", bundled.ResourceIsolatedThunderstormsSvg),
								lowhightempcard.View("Friday", "32°", "41°", bundled.ResourceIsolatedThunderstormsSvg),
								lowhightempcard.View("Saturday", "22°", "33°", bundled.ResourceIsolatedThunderstormsSvg),
								lowhightempcard.View("Sunday", "35°", "45°", bundled.ResourceIsolatedThunderstormsSvg),
							),
						),
					),
				),
			),
		),
	)
}
