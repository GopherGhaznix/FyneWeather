package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/GopherGhaznix/FyneWeather/components/headerscreen1"
	"github.com/GopherGhaznix/FyneWeather/components/hourlyDetails"
	"github.com/GopherGhaznix/FyneWeather/components/todayWeather"
	"github.com/GopherGhaznix/FyneWeather/components/weatherInfo"
)

func TodayScreen() fyne.CanvasObject {

	return container.NewScroll(
		container.NewPadded(
			container.NewPadded(
				container.NewBorder(
					headerscreen1.View(),
					nil, nil, nil, container.NewPadded(
						container.NewVBox(
							todayWeather.View(),
							canvas.NewText("", color.Transparent), // just for decent amount of spacing
							weatherInfo.View(),
							hourlyDetails.View(),
						),
					),
				),
			),
		),
	)
}
