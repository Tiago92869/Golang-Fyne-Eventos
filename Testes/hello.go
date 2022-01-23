package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	//DEFINE THE SIZE OF THE WINDOW (WIDTH, HEIGHT)
	w.Resize(fyne.NewSize(700, 400))
	w.SetContent(widget.NewLabel("Hello Fyne!"))

	w.ShowAndRun()

}
