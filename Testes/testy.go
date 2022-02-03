package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Golang-Eventos")
	window.Resize(fyne.NewSize(900, 500))

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(
			//o primeiro element e um icon
			//o segundo é a ação
			theme.MediaPlayIcon(), func() {
				fmt.Println("Buy Tickets")
			},
		),
		widget.NewToolbarAction(
			theme.CancelIcon(), func() {
				fmt.Println("Check Details")
			},
		),

		widget.NewToolbarAction(
			theme.HelpIcon(), func() {
				fmt.Println("help button")
			},
		),
	)
	c := container.NewHBox(toolbar)
	c.Move(fyne.NewPos(1, 52))
	window.SetContent(c)
	window.ShowAndRun()
}
