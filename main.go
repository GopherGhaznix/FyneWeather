package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	ui "github.com/GopherGhaznix/FyneWeather/UI"
	"github.com/GopherGhaznix/FyneWeather/bundled"
	"github.com/GopherGhaznix/FyneWeather/config"
	"github.com/GopherGhaznix/FyneWeather/state"
)

func init() {
	config.FontTekoBold = fyne.NewStaticResource("Teko Light", bundled.ResourceTekoSemiBoldTtf.Content())
	config.FontTekoNormal = fyne.NewStaticResource("Teko Light Normal", bundled.ResourceTekoRegularTtf.Content())
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(320, 580))

	// init boths screens
	state.TodayWindow = ui.TodayScreen()
	state.WeeklyWindow = ui.WeeklyScreen()
	state.CurrentWindow = ui.ScreenHandler()

	w.SetContent(state.CurrentWindow)
	w.ShowAndRun()
}
