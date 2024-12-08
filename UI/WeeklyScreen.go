package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
								lowhightempcard.View("Today", "./assets/animated/isolated-thunderstorms.svg", "18°", "21°"),
								lowhightempcard.View("Thesday", "./assets/animated/isolated-thunderstorms.svg", "13°", "24°"),
								lowhightempcard.View("Wednesday", "./assets/animated/isolated-thunderstorms.svg", "23°", "31°"),
								lowhightempcard.View("Thusday", "./assets/animated/isolated-thunderstorms.svg", "15°", "29°"),
								lowhightempcard.View("Friday", "./assets/animated/isolated-thunderstorms.svg", "32°", "41°"),
								lowhightempcard.View("Saturday", "./assets/animated/isolated-thunderstorms.svg", "22°", "33°"),
								lowhightempcard.View("Sunday", "./assets/animated/isolated-thunderstorms.svg", "35°", "45°"),
							),
						),
					),
				),
			),
		),
	)
}
