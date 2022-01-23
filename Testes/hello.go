package main

import (
	"fmt"
	"image/color"

	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//IMPORTING FYNE IN MY PROJECT
func main() {
	fmt.Println("Hello World")
	//start with go mod init myapp to create a package
	//we will createour frist fyne project

	a := app.New()                   //our first line of code will creating a new app
	w := a.NewWindow("Hello")        //we create a new window with a title
	w.Resize(fyne.NewSize(800, 500)) //DEFINE THE SIZE OF THE WINDOW (WIDTH, HEIGHT)

	//another way to insert labels
	labelX := widget.NewLabel("here is second label") //create a label
	labelX.Move(fyne.NewPos(100, 300))                //POSITION OF THE LABEL

	//create a button
	btn := widget.NewButton("button Name", func() {
		//when clicked prints a text in terminal
		fmt.Println("BUTTON CLICKED")
	})
	btn.Resize(fyne.NewSize(150, 30)) //Size of new button
	btn.Move(fyne.NewPos(100, 100))   //Position of the button

	//check box widget
	checkbox1 := widget.NewCheck("Male", func(b bool) {
		if b == true {
			labelX.Text = "Male"
			labelX.Refresh()
		} else {
			labelX.Text = "deselected"
			labelX.Refresh()
		}
	})

	//check box widget
	checkbox2 := widget.NewCheck("Female", func(b bool) {
		if b == true {
			labelX.Text = "Female"
			labelX.Refresh()
		} else {
			labelX.Text = "deselected"
			labelX.Refresh()
		}
	})

	//creating an url
	url, _ := url.Parse("https://google.com")

	//hyperlink Widget
	//first value is label
	//2nd value is url/ website address
	hyperlink := widget.NewHyperlink("visit me", url)

	//text, first value isnthe text, second is the colors
	//create colors (R: red, G: green, B: blue, A: opacity )
	colorx := color.NRGBA{R: 50, G: 255, B: 125, A: 255}
	textX := canvas.NewText("Here is my text", colorx)

	//IMAGE
	//imgame aparece buggada por ser NewVBox
	img := canvas.NewImageFromFile("toji.jpg")

	circle := canvas.NewCircle(color.NRGBA{R: 0, G: 244, B: 155, A: 250})
	circle.StrokeColor = color.NRGBA{R: 150, G: 0, B: 0, A: 250}
	circle.StrokeWidth = 33

	//	ICON
	Iconx := widget.NewIcon(theme.CancelIcon())

	//INSERIR OS ELEMENTOS NA JANELA
	w.SetContent(
		container.NewVBox(
			labelX,
			btn,
			checkbox1,
			checkbox2,
			hyperlink,
			textX,
			img,
			circle,
			Iconx,
		),
	)

	//w.SetContent(circle)
	w.ShowAndRun() //finally running our app
}
