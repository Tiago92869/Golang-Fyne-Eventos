package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("teste")
	w.Resize(fyne.NewSize(400, 400))
	w.SetFullScreen(true)

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter your name")
	entry.Resize(fyne.NewSize(250, 30))
	entry.Move(fyne.NewPos(40, 50))

	entry_email := widget.NewEntry()
	entry_email.SetPlaceHolder("Enter your email")
	entry_email.Resize(fyne.NewSize(250, 30))
	entry_email.Move(fyne.NewPos(40, 100))

	entry_address := widget.NewEntry()
	entry_address.SetPlaceHolder("Enter your name")
	entry_address.Resize(fyne.NewSize(250, 30))
	entry_address.Move(fyne.NewPos(40, 150))

	//button
	btn_submit := widget.NewButton("Submit", nil)
	btn_submit.Resize(fyne.NewSize(100, 30))
	btn_submit.Move(fyne.NewPos(40, 200))

	w.SetContent(

		container.NewWithoutLayout(
			entry,
			entry_email,
			entry_address,
			btn_submit,
		),
	) //show and run
	w.ShowAndRun()

}
