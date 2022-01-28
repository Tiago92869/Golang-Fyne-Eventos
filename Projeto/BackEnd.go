package main

import (
	"fmt"
	"time"
)

//STRUCT DATE FOR THE BEGGINING OF THE SHOW
type DataInicio struct {
	AnoI    int
	MesI    int
	DiaI    int
	HoraI   int
	MinutoI int
}

//STRUCT DATE FOR THE END OF THE SHOW
type DataFim struct {
	AnoF    int
	MesF    int
	DiaF    int
	HoraF   int
	MinutoF int
}

//STRUCT OF BILHETES
type Bilhete struct {
	ID int
}

//STRUCT OF EVENTO
type Evento struct {
	Nome          string
	DataInicio    DataInicio
	DataFim       DataFim
	Participantes int
	Bilhete       []Bilhete
	Preço         float64
}

//ARRAY OF ALL EVENTOS CREATED
var lista_eventos []Evento

//CREATE A NEW EVENTO
func registoEvento(nome string, anoi int, mesi int, diai int, horai int, mini int, anof int, mesf int, diaf int, horaf int, minf int, participantes int, preço float64) *Evento {
	//RETURN A NEW EVENTO
	return &Evento{
		Nome: nome,
		DataInicio: DataInicio{
			AnoI:    anoi,
			MesI:    mesi,
			DiaI:    diai,
			HoraI:   horai,
			MinutoI: mini,
		},
		DataFim: DataFim{
			AnoF:    anof,
			MesF:    mesf,
			DiaF:    diaf,
			HoraF:   horaf,
			MinutoF: minf,
		},
		Participantes: participantes,
		Preço:         preço,
	}
}

//ADD THE EVENTO CREATED TO THE ARRAY OF EVENTOS AND ALSO CRIATE THE SAME EVENTO USING THE FUNCTION registoEvento
func adicionarEvento(nome string, anoi int, mesi int, diai int, horai int, mini int, duração int, participantes int, preço float64) int {
	//FIRST WE CHECK IF THERES IS ALREADY A EVENTO WITH THE SAME NAME AS THE ONE WE WILL CREATE
	for i := 0; i < len(lista_eventos); i++ {
		//CHECK IF THE NAME ALREADY EXISTS IN THE ARRAY OF EVENTOS
		if nome == lista_eventos[i].Nome {
			//IF THERES ALREADY THE NAME THE SYSTEM RETURNS A MESSAGE
			println("O nome introduzido já existe")
			return 1
		}
	}
	//NOW WE CALCULATE THE DATE OF THE END OF THE EVENTO WITH THE DURATION GIVEN
	timei := time.Date(anoi, time.Month(mesi), diai, horai, mini, 0, 0, time.UTC)
	//WE ADD TO THE DATE OF THE BEGGINING OFTHE SHOW THE DURATION AND GET THE DATE OF THE END OF THE SHOW
	newtime := timei.Add(time.Hour * time.Duration(duração))
	//GET OF EACH DATE WE NEED
	var anof int = newtime.Year()
	//THE MONTH NEEDS TO BE CONVERTED BECAUSE ITS NOT A INT
	var m = newtime.Month()
	var mesf int = int(m)
	var diaf int = newtime.Day()
	var horaf int = newtime.Hour()
	var minf int = newtime.Minute()
	//WE NOW CHECK IF THERE IS ALREADY A EVENTO AT THE SAME TIME
	//FRIST WE CREATE AN ARRAY TO SAVE THE EVENTOS THAT ARE APPENING AT THE SAME TIME
	var aa []Evento
	for i := 0; i < len(lista_eventos); i++ {
		//CREATE A DATE INICIAL OF THE EVENTO IN THE ARRAY
		datei := time.Date(lista_eventos[i].DataInicio.AnoI, time.Month(lista_eventos[i].DataInicio.MesI), lista_eventos[i].DataInicio.DiaI, lista_eventos[i].DataInicio.HoraI, lista_eventos[i].DataInicio.MinutoI, 0, 0, time.UTC)
		//CREATE THE END DATE FOT THE EVENTO IN THE ARRAY
		datef := time.Date(lista_eventos[i].DataFim.AnoF, time.Month(lista_eventos[i].DataFim.MesF), lista_eventos[i].DataFim.DiaF, lista_eventos[i].DataFim.HoraF, lista_eventos[i].DataFim.MinutoF, 0, 0, time.UTC)
		//NOW WE COMPARE DATES SO THAT THE EVENTO WE WANT TO HAD DOESN'T ENTER IN CONFLICT WITH OTHER EVENTO ALREADY CREATED
		if ((datei == timei) || (datei == newtime) || (datef == timei) || (datef == newtime)) || ((timei.After(datei)) && (timei.Before(datef))) || ((newtime.After(datei)) && (newtime.Before(datef))) {
			//WE ADD THE EVENTO THAT HAVE CONFLICTED DATES TO AN ARRAY
			aa = append(aa, lista_eventos[i])
		}
	}
	if len(aa) > 0 {
		//IF THE ARRAY IS LONGER THAN 0 IT SHALL PRINT THE EVENTOS WITH CONFLICTED DATES AND CLOSE THE FUNCTION
		fmt.Println("JÁ EXISTE DATAS MARCADAS: \n", aa)
		return 1
	}
	//CREATE A NEW EVENTO
	registoEvento(nome, anoi, mesi, diai, horai, mini, anof, mesf, diaf, horaf, minf, participantes, preço)
	//ADD THE EVENTO TO THE ARRAY LIST
	lista_eventos = append(lista_eventos, *registoEvento(nome, anoi, mesi, diai, horai, mini, anof, mesf, diaf, horaf, minf, participantes, preço))
	return 0
}

//GET OF THE ARRAY OF EVENTOS
func GetEventosList() []Evento {
	return lista_eventos
}

//GET OF THE EVENTOS THAT ALREADY HAPPENED
func GetEventosHappened() []Evento {
	//CREATE THE LIST
	var lista []Evento
	//CHECK THE ARRAY OF ALL EVENTOS
	for i := 0; i < len(lista_eventos); i++ {
		//CREATE THE END DATE FOT THE EVENTO IN THE ARRAY
		datef := time.Date(lista_eventos[i].DataFim.AnoF, time.Month(lista_eventos[i].DataFim.MesF), lista_eventos[i].DataFim.DiaF, lista_eventos[i].DataFim.HoraF, lista_eventos[i].DataFim.MinutoF, 0, 0, time.UTC)
		//GET THE CURRENT DATE
		currentTime := time.Now()
		//IF THE DATE OF THE EVENTO IS BEFORE THE CURRENT DATE ADD TO THE ARRAY
		if datef.Before(currentTime) {
			//ADD THE EVENTO TO THE ARRAY
			lista = append(lista, lista_eventos[i])
		}
	}
	//RETURN OF THE ARRAY
	return lista
}

//DELETE AN EVENTO FROM THE ARRAY
func deleteEvento(nome string) {
	//SEARCH THE ENTIRE ARRAY TO SEARCH FOR THE EVENTO WITH THE SAME NAME
	for i := 0; i < len(lista_eventos); i++ {
		//IF THE EVENTO IS FOUND DELETE IT
		if lista_eventos[i].Nome == nome {
			//append(value to replace, replace value...-> till the end)
			lista_eventos = append(lista_eventos[:i], lista_eventos[i+1:]...)
		}
	}
}

//SEARCH AN EVENTO FROM THE ARRAY
func getEvento(nome string) Evento {
	var a Evento
	//SEARCH THE ENTIRE ARRAY TO SEARCH FOR THE EVENTO WITH THE SAME NAME
	for i := 0; i < len(lista_eventos); i++ {
		//IF THE EVENTO IS FOUND RETURN IT
		if lista_eventos[i].Nome == nome {
			a = lista_eventos[i]
		}
	}
	return a
}

//EDIT AN SPECIFIC EVENTO

func main() {
	adicionarEvento("aaa", 1, 3, 4, 8, 6, 3, 23, 3)
	adicionarEvento("aafa", 1, 3, 4, 20, 6, 2, 23, 3)
	adicionarEvento("aafaaaa", 5551, 3, 5, 20, 6, 2, 23, 3)
	fmt.Println(GetEventosList())
	println(len(lista_eventos))
	getEvento("aafa")
	fmt.Println(getEvento("aafa"))
	fmt.Println(GetEventosHappened())
	println(len(lista_eventos[0].Bilhete))

}
