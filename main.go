package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	ui "github.com/GopherGhaznix/FyneWeather/UI"
	"github.com/GopherGhaznix/FyneWeather/config"
	"github.com/GopherGhaznix/FyneWeather/state"
)

func init() {
	fontBytes, err := os.ReadFile("./assets/font/Teko/static/Teko-SemiBold.ttf")
	if err != nil {
		log.Fatal(err)
	}
	config.FontTekoBold = fyne.NewStaticResource("Teko Light", fontBytes)

	font2Bytes, err := os.ReadFile("./assets/font/Teko/static/Teko-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	config.FontTekoNormal = fyne.NewStaticResource("Teko Light Normal", font2Bytes)
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
