package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"log"
)

type (
	application struct {
		fyne.App
	}
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	a := &application{App: app.New()}
	a.Settings().SetTheme(theme.LightTheme())
	win := a.NewWindow("Timers")
	hour := NewNumericFieldRange(0, 999)
	min := NewNumericFieldRange(0, 60)
	min.Text = "10"
	sec := NewNumericFieldRange(0, 60)

	button1 := widget.NewButton("GO", nil)

	win.SetContent(
		widget.NewVBox(
			widget.NewHBox(
				hour, min, sec,
				button1,
			),
			widget.NewHBox(
				widget.NewButton("+", func() {}),
				widget.NewButton("Quit", func() { a.Quit() }),
			),
		),
	)
	win.ShowAndRun()
	return nil
}

//func main() {
//	a := app.New()
//	win := a.NewWindow("Window Test")
//	hour := widget.NewEntry()
//	hour.Text = "00"
//	hour.OnChanged = hourChanged
//	min := widget.NewEntry()
//	min.Text = "14"
//	sec := widget.NewEntry()
//	sec.Text = "59"
//	win.SetContent(
//		widget.NewVBox(
//			widget.NewHBox(
//				hour, min, sec,
//				widget.NewButton("GO", func() { showTime(hour.Text, min.Text, sec.Text) }),
//			),
//			widget.NewHBox(
//				widget.NewButton("+", func() {}),
//				widget.NewButton("Quit", func() { a.Quit() }),
//			),
//		),
//	)
//	win.ShowAndRun()
//}

//func main() {
//	a := &application{App: app.New()}
//	win := a.NewWindow("Timer")
//	win.SetContent(
//		widget.NewVBox(
//			widget.NewLabel("Click a button below to run a test"),
//			widget.NewButton("Progress Bar", a.testProgressBar),
//			widget.NewButton("Email", a.testEmail),
//			widget.NewButton("Timer", a.testTimer),
//			widget.NewCheck("check", func(b bool){
//				print(b)
//			}),
//			widget.NewEntry(),
//			widget.NewMultiLineEntry(),
//			widget.NewProgressBar(),
//			widget.NewButton("Quit", func() { a.Quit() }),
//		),
//	)
//	win.ShowAndRun()
//}
