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
	w.Resize(´fynefyne.MeasureText())
	w.SetFullScreen(false)

	fmt.Println("Testando o Go - xiko")

	//entry vai ser uma variavel (logica da struct/ em que em vez de termos outras variaveis dentro/ vamos definir as suas propriedades)
	nome := widget.NewEntry()
	nome.SetPlaceHolder("Aqui fica o txt do Label(nome)")
	nome.Resize(fyne.NewSize(200, 40)) //definir as medidas da box (pre-defnido)
	nome.Move(fyne.NewPos(40,50)) //Definir a translação da box nas medidas que queremos

	




}