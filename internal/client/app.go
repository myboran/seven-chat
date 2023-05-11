package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func NewClient() {
	myApp := app.New()
	w := myApp.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	rect := canvas.NewRectangle(color.Black)
	w.SetContent(rect)
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
