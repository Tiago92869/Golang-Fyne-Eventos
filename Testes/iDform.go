package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

import "fmt"

func main(){

	a := app.New();
	w := a.NewWindow("tete")
	//w.Resize(´fynefyne.MeasureText())
	w.SetFullScreen(true)

	fmt.Println("Testando o Go - xiko")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Aqui fica o txt do Label(nome)")
	entry.Resize(fyne.NewSize(200, 40)) //definir as medidas da box (pre-defnido)
	entry.Move(fyne.NewPos(40,50)) //Definir a translação da box nas medidas que queremos




}