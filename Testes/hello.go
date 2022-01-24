package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

//IMPORTING FYNE IN MY PROJECT
func main() {
	fmt.Println("O xiko esteve aqui")
	//start with go mod init myapp to create a package
	//we will create our first fyne project

	a := app.New() //our first line of code will creating a new app

	w := a.NewWindow("Hello") //we create a new window with a title

	w.Resize(fyne.NewSize(800, 500)) //DEFINE THE SIZE OF THE WINDOW (WIDTH, HEIGHT)

	w.SetContent(widget.NewLabel("Hello Fyne!")) //Insert label to the window

	//another way to insert labels
	labelX := widget.NewLabel("here is second label") //create a label
	w.SetContent(labelX)                              //insert label in window

	//create a button
	btn := widget.NewButton("button Name", func() {
		//When clicked it will print this
		fmt.Println("BUTTON HAS BEEN CLICKED")
	})

	btn.Resize(fyne.NewSize(150, 30)) //Size of new button
	btn.Move(fyne.NewPos(200, 200))   //Position of the button

	w.SetContent(btn) //add the button to the window

	w.ShowAndRun() //finally running our app

}

func Mark() {
	panic("unimplemented")
}
