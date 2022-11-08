package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rainu/go-calendar/internal/calendar"
	"os"
	"time"
)

type mainFrame struct {
	app    fyne.App
	window fyne.Window

	currentDay   time.Time
	currentMonth calendar.Month

	contentContainer *fyne.Container
	btnMonth         *widget.Button
}

func NewMainFrame(t time.Time) *mainFrame {
	result := &mainFrame{
		currentDay:   t,
		currentMonth: calendar.NewMonth(t),
	}

	result.app = app.New()

	if dd, ok := result.app.Driver().(desktop.Driver); ok {
		result.window = dd.CreateSplashWindow()
	} else {
		result.window = result.app.NewWindow("Calendar")
	}

	result.window.SetMaster()
	result.window.SetFixedSize(true)
	result.app.Settings().SetTheme(customTheme{result.app.Settings().Theme()})

	result.btnMonth = widget.NewButton("", result.onButtonMonthClick)
	result.btnMonth.Alignment = widget.ButtonAlignCenter
	result.contentContainer = container.New(layout.NewGridLayout(8))

	result.window.SetContent(container.NewVBox(
		container.New(
			layout.NewGridLayout(3),
			widget.NewButtonWithIcon("", theme.NavigateBackIcon(), result.setToPreviousMonth),
			result.btnMonth,
			widget.NewButtonWithIcon("", theme.NavigateNextIcon(), result.setToNextMonth),
		),
		result.contentContainer,
	))

	result.window.Canvas().SetOnTypedKey(result.onKeyEvent)

	return result
}

func (f *mainFrame) setToPreviousMonth() {
	t := time.Date(
		f.currentDay.Year(),
		f.currentDay.Month()-1,
		f.currentDay.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	f.setCurrentMonth(t)
}

func (f *mainFrame) setToPreviousYear() {
	t := time.Date(
		f.currentDay.Year()-1,
		f.currentDay.Month(),
		f.currentDay.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	f.setCurrentMonth(t)
}

func (f *mainFrame) setToNextMonth() {
	t := time.Date(
		f.currentDay.Year(),
		f.currentDay.Month()+1,
		f.currentDay.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	f.setCurrentMonth(t)
}

func (f *mainFrame) setToNextYear() {
	t := time.Date(
		f.currentDay.Year()+1,
		f.currentDay.Month(),
		f.currentDay.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	f.setCurrentMonth(t)
}

func (f *mainFrame) onButtonMonthClick() {
	f.setCurrentMonth(time.Now())
}

func (f *mainFrame) onKeyEvent(e *fyne.KeyEvent) {
	switch e.Name {
	case "Right", "D":
		f.setToNextMonth()
	case "Left", "A":
		f.setToPreviousMonth()
	case "Up", "W":
		f.setToNextYear()
	case "Down", "S":
		f.setToPreviousYear()
	case "Space":
		f.setCurrentMonth(time.Now())
	case "Escape":
		os.Exit(0)
	}
}

func (f *mainFrame) setCurrentMonth(t time.Time) {
	f.currentDay = t
	f.currentMonth = calendar.NewMonth(t)
	f.update()
}

func (f *mainFrame) update() {
	f.contentContainer.RemoveAll()

	f.btnMonth.SetText(fmt.Sprintf("%s %d", MonthNames[f.currentDay.Month()], f.currentDay.Year()))
	for i, _ := range f.currentMonth {
		for j, _ := range f.currentMonth[i] {
			switch v := f.currentMonth[i][j].(type) {
			case calendar.Day:
				if v.IsEdge() {
					txt := canvas.NewText(
						fmt.Sprintf("%d", v.Day()),
						f.app.Settings().Theme().Color(
							theme.ColorNameHover,
							f.app.Settings().ThemeVariant(),
						),
					)
					txt.Alignment = fyne.TextAlignCenter
					if v.IsSpecial() {
						txt.Color = f.app.Settings().Theme().Color(
							theme.ColorNamePlaceHolder,
							f.app.Settings().ThemeVariant(),
						)
					}

					f.contentContainer.Add(txt)
				} else if v.IsCurrent() {
					btn := widget.NewButton(fmt.Sprintf("%d", v.Day()), nil)
					if v.IsSpecial() {
						btn.Importance = widget.HighImportance
					}

					f.contentContainer.Add(btn)
				} else {
					txt := canvas.NewText(
						fmt.Sprintf("%d", v.Day()),
						f.app.Settings().Theme().Color(
							theme.ColorNameForeground,
							f.app.Settings().ThemeVariant(),
						),
					)
					txt.Alignment = fyne.TextAlignCenter

					if v.IsSpecial() {
						txt.Color = f.app.Settings().Theme().Color(
							theme.ColorNamePrimary,
							f.app.Settings().ThemeVariant(),
						)
					}

					f.contentContainer.Add(txt)
				}
			case time.Weekday:
				lbl := widget.NewLabel(DayNames[v][:2])
				lbl.TextStyle.Bold = true
				lbl.Alignment = fyne.TextAlignCenter

				f.contentContainer.Add(lbl)
			case calendar.CalendarWeek:
				lbl := widget.NewLabel(fmt.Sprintf("%d", v))
				lbl.TextStyle.Bold = true
				lbl.Alignment = fyne.TextAlignCenter

				f.contentContainer.Add(lbl)
			default:
				f.contentContainer.Add(&widget.BaseWidget{})
			}
		}
	}
	f.contentContainer.Refresh()
}

func (f *mainFrame) ShowAndRun() {
	f.update()
	f.window.ShowAndRun()
}
