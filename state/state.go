package state

import "fyne.io/fyne/v2"

var (
	CurrentWindow fyne.CanvasObject
	TodayWindow   fyne.CanvasObject
	WeeklyWindow  fyne.CanvasObject
)

func SwitchToday() {
	TodayWindow.Show()
	WeeklyWindow.Hide()

	CurrentWindow.Refresh()
	TodayWindow.Refresh()
	WeeklyWindow.Refresh()
}

func SwitchWeekly() {
	TodayWindow.Hide()
	WeeklyWindow.Show()

	CurrentWindow.Refresh()
	TodayWindow.Refresh()
	WeeklyWindow.Refresh()
}
