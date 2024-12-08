package weatherInfo

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/GopherGhaznix/FyneWeather/bundled"
	"github.com/GopherGhaznix/FyneWeather/components/infobox"
)

func View() fyne.CanvasObject {

	BG := canvas.NewRectangle(color.RGBA{R: 32, G: 35, B: 41, A: 255})
	BG.CornerRadius = 15

	return container.NewStack(
		BG,
		container.NewAdaptiveGrid(
			3,
			infobox.View(bundled.ResourceWindSvg, "Wind", "10 m/s"),
			infobox.View(bundled.ResourceSweatDropletsSvgrepoComSvg, "Humidity", "98%"),
			infobox.View(bundled.ResourceUmbrellaWithRainDropsSvg, "Rain", "100%"),
		),
	)
}
