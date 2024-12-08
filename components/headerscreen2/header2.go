package headerscreen2

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
	locText.TextSize = 20

	return container.NewBorder(
		nil,
		nil,
		widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {

			state.SwitchToday()
		}),
		widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {}),
		container.NewCenter(locText),
	)
}
