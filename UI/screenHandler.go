package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/GopherGhaznix/FyneWeather/state"
)

var ()

func ScreenHandler() fyne.CanvasObject {

	state.TodayWindow.Show()
	state.WeeklyWindow.Hide()

	return container.NewStack(
		state.TodayWindow,
		state.WeeklyWindow,
	)
}
