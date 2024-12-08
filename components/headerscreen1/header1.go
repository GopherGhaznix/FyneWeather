package headerscreen1

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/GopherGhaznix/FyneWeather/config"
	"github.com/GopherGhaznix/FyneWeather/state"
)

func View() fyne.CanvasObject {

	locText := canvas.NewText("Chakwal", theme.Color(theme.ColorNameForeground))
	locText.FontSource = config.FontTekoBold
	locText.TextSize = 25

	dateText := canvas.NewText("7 December, Sunday", theme.Color(theme.ColorNamePressed))
	dateText.TextSize = 12

	return container.NewBorder(
		nil, nil,
		container.NewVBox(
			locText,
			dateText,
		),
		container.NewCenter(
			widget.NewButtonWithIcon("", theme.GridIcon(), func() {

				state.SwitchWeekly()
			}),
		),
	)
}
