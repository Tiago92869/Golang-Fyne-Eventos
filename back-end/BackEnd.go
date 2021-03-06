package back

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	count         int
	Preço         float64
}

//ARRAY OF ALL EVENTOS CREATED
var Lista_eventos []Evento = LoadListaEventos()

//CREATE A NEW TICKET
func RegistoTicket(id int) *Bilhete {
	//RETURN A NEW TICKET
	return &Bilhete{
		ID: id,
	}
}

//CREATE A FILE EVENTOS IF IT DOENS'T EXIST
func CriarPastaEcentos() {
	_, ferr := os.Open("eventos.txt")
	if ferr != nil {
		file, err := os.Create("eventos.txt")
		if err != nil {
			log.Fatal("We got error", err)
		}
		defer file.Close()
	}
}

//FUNCTION TO SAVE THE TICKET IN THE EVENTO FROM THE TEXT FILE
func AdicionarBilhete(nome string, id int, listaE []Evento) {
	for i := 0; i < len(listaE); i++ {
		if listaE[i].Nome == nome {
			RegistoTicket(id)
			listaE[i].Bilhete = append(listaE[i].Bilhete, *RegistoTicket(id))
		}
	}
}

//CREATES THE FILE AND SAVES DE DATA TO THE FILE
func GuardarListaEventos() {
	CriarPastaEcentos()
	file2, _ := os.Create("eventos.txt")
	for i := 0; i < len(Lista_eventos); i++ {
		//SAVE NAME
		nome := (Lista_eventos[i].Nome) + ","
		//WRITE IN THE FILE THE NAME OF EVENTO
		file2.WriteString(nome)
		//SAVE BEGINNING YEAR
		anoi := strconv.Itoa(Lista_eventos[i].DataInicio.AnoI) + ","
		//WRITE IN THE FILE THE BEGGINING YEAR
		file2.WriteString(anoi)
		//SAVE BEGINNING MONTH
		mesi := strconv.Itoa(Lista_eventos[i].DataInicio.MesI) + ","
		//WRITE IN FILE THE BEGGINING MONTH
		file2.WriteString(mesi)
		//SAVE BEGINNING DAY
		diai := strconv.Itoa(Lista_eventos[i].DataInicio.DiaI) + ","
		//WRITE IN FILE THE BEGINNING DAY
		file2.WriteString(diai)
		//SAVE BEGINNING HOUR
		horai := strconv.Itoa(Lista_eventos[i].DataInicio.HoraI) + ","
		//WRITE IN THE FILE THE BEGINNING HOUR
		file2.WriteString(horai)
		//SAVE BEGINNING MIN
		mini := strconv.Itoa(Lista_eventos[i].DataInicio.MinutoI) + ","
		//WRITE IN THE FILE THE BEGINNING MIN
		file2.WriteString(mini)
		//SAVE END YEAR
		anof := strconv.Itoa(Lista_eventos[i].DataFim.AnoF) + ","
		//WRITE IN THE FILE THE END YEAR
		file2.WriteString(anof)
		//SAVE END MONTH
		mesf := strconv.Itoa(Lista_eventos[i].DataFim.MesF) + ","
		//WRITE IN THE FILE THE END MONTH
		file2.WriteString(mesf)
		//SAVE END DAY
		diaf := strconv.Itoa(Lista_eventos[i].DataFim.DiaF) + ","
		//WRITE IN THE FILE THE END DAY
		file2.WriteString(diaf)
		//SAVE END HOUR
		horaf := strconv.Itoa(Lista_eventos[i].DataFim.HoraF) + ","
		//WRITE IN THE FILE THE END HOUR
		file2.WriteString(horaf)
		//SAVE END MIN
		minf := strconv.Itoa(Lista_eventos[i].DataFim.MinutoF) + ","
		//WRITE IN THE FILE THE END MINUTE
		file2.WriteString(minf)
		//SAVE THE NUMBER OF PARTICIPANTES
		partic := strconv.Itoa(Lista_eventos[i].Participantes) + ","
		//WRITE IN THE FILE THE NUMBER OF PARTICIPANTES
		file2.WriteString(partic)
		//SAVE THE COUNT OF TICKETS
		count := strconv.Itoa(Lista_eventos[i].count) + ","
		//WRITE IN THE FILE THE COUNT OF TICKETS
		file2.WriteString(count)
		//SAVE THE PRICE PER TICKET
		preço := fmt.Sprintf("%.2f", Lista_eventos[i].Preço)
		//WRITE IN THE FILE THE PRICE PER TICKET
		file2.WriteString(preço)
		if len(Lista_eventos[i].Bilhete) == 0 {
			//CREATES A NEW LINE
			file2.WriteString("\n")
		} else {
			//ELSE ADD COMON TO ADD A NEW TICKET
			file2.WriteString(",")
			//WE GO TO EACH TICKET FROM THE CURRENT EVENTO
			for ii := 0; ii < len(Lista_eventos[i].Bilhete); ii++ {
				var bilhete string
				//IF ITS THE LAST TICKET WE MAKE A NEW LINE
				if ii == (len(Lista_eventos[i].Bilhete))-1 {
					//SAVE THE ID OF THE TICKET
					bilhete = strconv.Itoa(Lista_eventos[i].Bilhete[ii].ID) + "\n"
					//WRITE IN THE FILE THE ID OF THE TICKET
					file2.WriteString(bilhete)
				} else {
					//SAVE THE ID OF THE TICKET
					bilhete = strconv.Itoa(Lista_eventos[i].Bilhete[ii].ID) + ","
					//WRTIE IN THE FILE THE ID OF THE TICKET
					file2.WriteString(bilhete)
				}
			}
		}
	}
}

//LOAD THE EVENTOS FROM FILE TO ARRAY OF EVENTOS
func LoadListaEventos() []Evento {
	CriarPastaEcentos()
	var list []Evento
	file, err := os.Open("eventos.txt")
	if err != nil {
		log.Fatal("We got error", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		//SCANNER FOR EACH LINE
		line := fileScanner.Text()
		//SEPARATE THE ELEMENTS BY ","
		items := strings.Split(line, ",")
		//SAVE THE VALUES OF THE FILE
		//GET NAME
		nome := items[0]
		//GET BEGINNING YEAR
		anoi, _ := strconv.Atoi(items[1])
		//GET BEGINNING MOUNTH
		mesi, _ := strconv.Atoi(items[2])
		//GET BEGINNING DAY
		diai, _ := strconv.Atoi(items[3])
		//GET BEGINNING HOUR
		horai, _ := strconv.Atoi(items[4])
		//GET BEGINNING MIN
		mini, _ := strconv.Atoi(items[5])
		//GET END YEAR
		anof, _ := strconv.Atoi(items[6])
		//GET END MOUNTH
		mesf, _ := strconv.Atoi(items[7])
		//GET END DAY
		diaf, _ := strconv.Atoi(items[8])
		//GET END HOUR
		horaf, _ := strconv.Atoi(items[9])
		//GET END MIN
		minf, _ := strconv.Atoi(items[10])
		//GET NUMBER OF PARTICIPANTES
		partic, _ := strconv.Atoi(items[11])
		//GET COUNT OF ID TICKETS
		count, _ := strconv.Atoi(items[12])
		//GET PRICE PER TICKET
		preço, _ := strconv.ParseFloat(items[13], 64)
		//INSERT INTO A LIST
		list = append(list, *RegistoEvento(nome, anoi, mesi, diai, horai, mini, anof, mesf, diaf, horaf, minf, partic, count, preço))
		//WE CHECK IF THERE ARE MORE ITENS IN THIS CASE ARE TICKETS
		if len(items) > 13 {
			//IF THERE IS TICKETS WE ADD THEM TO THE EVENTO IN "list"
			for i := 14; i < len(items); i++ {
				id, _ := strconv.Atoi(items[i])
				AdicionarBilhete(nome, id, list)
			}
		}
	}
	//RETURN THE LIST
	return list
}

//CREATE A NEW EVENTO
func RegistoEvento(nome string, anoi int, mesi int, diai int, horai int, mini int, anof int, mesf int, diaf int, horaf int, minf int, participantes int, count int, preço float64) *Evento {
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
		count:         count,
		Preço:         preço,
	}
}

//ADD THE EVENTO CREATED TO THE ARRAY OF EVENTOS AND ALSO CRIATE THE SAME EVENTO USING THE FUNCTION registoEvento
func AdicionarEvento(nome string, anoi int, mesi int, diai int, horai int, mini int, duração int, participantes int, count int, preço float64) int {
	//FIRST WE CHECK IF THERES IS ALREADY A EVENTO WITH THE SAME NAME AS THE ONE WE WILL CREATE
	for i := 0; i < len(Lista_eventos); i++ {
		//CHECK IF THE NAME ALREADY EXISTS IN THE ARRAY OF EVENTOS
		if nome == Lista_eventos[i].Nome {
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
	for i := 0; i < len(Lista_eventos); i++ {
		//CREATE A DATE INICIAL OF THE EVENTO IN THE ARRAY
		datei := time.Date(Lista_eventos[i].DataInicio.AnoI, time.Month(Lista_eventos[i].DataInicio.MesI), Lista_eventos[i].DataInicio.DiaI, Lista_eventos[i].DataInicio.HoraI, Lista_eventos[i].DataInicio.MinutoI, 0, 0, time.UTC)
		//CREATE THE END DATE FOT THE EVENTO IN THE ARRAY
		datef := time.Date(Lista_eventos[i].DataFim.AnoF, time.Month(Lista_eventos[i].DataFim.MesF), Lista_eventos[i].DataFim.DiaF, Lista_eventos[i].DataFim.HoraF, Lista_eventos[i].DataFim.MinutoF, 0, 0, time.UTC)
		//NOW WE COMPARE DATES SO THAT THE EVENTO WE WANT TO HAD DOESN'T ENTER IN CONFLICT WITH OTHER EVENTO ALREADY CREATED
		if ((datei == timei) || (datei == newtime) || (datef == timei) || (datef == newtime)) || ((timei.After(datei)) && (timei.Before(datef))) || ((newtime.After(datei)) && (newtime.Before(datef))) {
			//WE ADD THE EVENTO THAT HAVE CONFLICTED DATES TO AN ARRAY
			aa = append(aa, Lista_eventos[i])
		}
	}
	if len(aa) > 0 {
		//IF THE ARRAY IS LONGER THAN 0 IT SHALL PRINT THE EVENTOS WITH CONFLICTED DATES AND CLOSE THE FUNCTION
		fmt.Println("JÁ EXISTE DATAS MARCADAS: \n", aa)
		return 1
	}
	//CREATE A NEW EVENTO
	RegistoEvento(nome, anoi, mesi, diai, horai, mini, anof, mesf, diaf, horaf, minf, participantes, count, preço)
	//ADD THE EVENTO TO THE ARRAY LIST
	Lista_eventos = append(Lista_eventos, *RegistoEvento(nome, anoi, mesi, diai, horai, mini, anof, mesf, diaf, horaf, minf, participantes, count, preço))
	GuardarListaEventos()
	return 0
}

//GET OF THE ARRAY OF EVENTOS
func GetEventosList() []Evento {
	return Lista_eventos
}

//GET OF THE EVENTOS THAT ALREADY HAPPENED
func GetEventosHappened() []Evento {
	//CREATE THE LIST
	var lista []Evento
	//CHECK THE ARRAY OF ALL EVENTOS
	for i := 0; i < len(Lista_eventos); i++ {
		//CREATE THE END DATE FOT THE EVENTO IN THE ARRAY
		datef := time.Date(Lista_eventos[i].DataFim.AnoF, time.Month(Lista_eventos[i].DataFim.MesF), Lista_eventos[i].DataFim.DiaF, Lista_eventos[i].DataFim.HoraF, Lista_eventos[i].DataFim.MinutoF, 0, 0, time.UTC)
		//GET THE CURRENT DATE
		currentTime := time.Now()
		//IF THE DATE OF THE EVENTO IS BEFORE THE CURRENT DATE ADD TO THE ARRAY
		if datef.Before(currentTime) {
			//ADD THE EVENTO TO THE ARRAY
			lista = append(lista, Lista_eventos[i])
		}
	}
	//RETURN OF THE ARRAY
	return lista
}

//DELETE AN EVENTO FROM THE ARRAY
func DeleteEvento(nome string) {
	//SEARCH THE ENTIRE ARRAY TO SEARCH FOR THE EVENTO WITH THE SAME NAME
	for i := 0; i < len(Lista_eventos); i++ {
		//IF THE EVENTO IS FOUND DELETE IT
		if Lista_eventos[i].Nome == nome {
			//append(value to replace, replace value...-> till the end)
			Lista_eventos = append(Lista_eventos[:i], Lista_eventos[i+1:]...)
		}
	}
	GuardarListaEventos()
}

//SEARCH AN EVENTO FROM THE ARRAY
func GetEvento(nome string) Evento {
	var a Evento
	//SEARCH THE ENTIRE ARRAY TO SEARCH FOR THE EVENTO WITH THE SAME NAME
	for i := 0; i < len(Lista_eventos); i++ {
		//IF THE EVENTO IS FOUND RETURN IT
		if Lista_eventos[i].Nome == nome {
			a = Lista_eventos[i]
		}
	}
	return a
}

//EDIT AN SPECIFIC EVENTO
func EditEvento(nome string, Nnome string, Nparticipantes int) int {
	//SEARCH THE ENTIRE ARRAY TO SEARCH FOR THE EVENTO WITH THE SAME NAME
	for i := 0; i < len(Lista_eventos); i++ {
		//IF THE EVENTO IS FOUND RETURN IT
		if Lista_eventos[i].Nome == nome {
			//CHANGE THE NAME
			Lista_eventos[i].Nome = Nnome
			//CHANGE THE NUMBER OF PEOPLE
			Lista_eventos[i].Participantes = Nparticipantes
		}
	}
	GuardarListaEventos()
	return 0
}

//BUY A TICKET
func BuyTicket(nome string) int {
	//WE SEARCH THE ENTIRE ARRAY FOR THE EVENTO WE WANT
	for i := 0; i < len(Lista_eventos); i++ {
		//WE COMPARE THE NAME OF THE EVENTO
		if Lista_eventos[i].Nome == nome {
			if len(Lista_eventos[i].Bilhete) >= Lista_eventos[i].Participantes {
				return 1
			} else {
				//WE GET THE COUNT OF THE EVENTO
				//THE COUNT REPRESENTES THE ID OF THE TICKET
				id := Lista_eventos[i].count
				//ADD ONE MORE TO COUNT
				Lista_eventos[i].count++
				//CREATE THE TICKET
				RegistoTicket(id)
				//ADD THE TICKET TO THE LIST OF TICKETS OF THE ARRAY
				Lista_eventos[i].Bilhete = append(Lista_eventos[i].Bilhete, *RegistoTicket(id))
			}
		}
	}
	GuardarListaEventos()
	//SUCCESS
	return 0
}

//RETURN A TICKET ALREADY BOUGHT
func ReturnTicket(nome string, id_ticket int) {
	//SEARCH THE ENTIRE ARRAY TO SEARCH FOR THE EVENTO WITH THE SAME NAME
	for i := 0; i < len(Lista_eventos); i++ {
		//IF THE EVENTO IS FOUND WE SEARCH HIS LIST OF TICKETS
		if Lista_eventos[i].Nome == nome {
			//SEARCH THE LIST OF TICKETS OF THE ARRAY
			for ii := 0; ii < len(Lista_eventos[i].Bilhete); ii++ {
				//IF THE TICKET ID IS THE SAME AS ABOVE DELETE
				if Lista_eventos[i].Bilhete[ii].ID == id_ticket {
					//REMOVE THE SPECIFIC TICKET FROM THE ARRAY OF TICKETS
					Lista_eventos[i].Bilhete = append(Lista_eventos[i].Bilhete[:ii], Lista_eventos[i].Bilhete[ii+1:]...)
				}
			}
		}
	}
	GuardarListaEventos()
}

//CHANGE THE PRICE OF AN EVENTO IF NO ONE BOUGHT THE TICKET
func ChangeTicketPrice(nome string, preço float64) int {
	//SEARCH THE ARRAY OF EVENTOS
	for i := 0; i < len(Lista_eventos); i++ {
		//CHECK FOR THE EVENTO WITH THE NAME WE WANT
		if Lista_eventos[i].Nome == nome {
			//CHECK IF THE ARE NO TICKETS SOLD
			if len(Lista_eventos[i].Bilhete) == 0 {
				//CHANGE THE PRICE OF THE TICKETS
				Lista_eventos[i].Preço = preço
			} else {
				//RETURN FAILURE
				return 1
			}
		}
	}
	GuardarListaEventos()
	//RETURN SUCCESS
	return 0
}

//COUNT ALL THE TICKETS
func CountAllTickets() int {
	//CREATE ALL COUNT FOR ALL THE TICKETS
	count_tickets := 0
	//WE SEARCH ALL THE ARRAY OF EVENTOS
	for i := 0; i < len(Lista_eventos); i++ {
		//COUNT THE NUMBER OF TICKETS PER EVENTO
		count_tickets += len(Lista_eventos[i].Bilhete)
	}
	//RETURN THE COUNT
	return count_tickets
}

/*
//RETURN DAY WITH THE MOST PEOPLE
func dayMostPeople() (int, int, int) {
	daywithmore := 0
	total := 0
	time := time.Now()
	for i := 0; i < len(Lista_eventos); i++ {
		if Lista_eventos[i].DataInicio.MesI == int(time.Month()) && Lista_eventos[i].DataFim.MesF == int(time.Month()) && Lista_eventos[i].DataInicio.AnoI == time.Year() && Lista_eventos[i].DataFim.AnoF == time.Year() {
			for ii := 0; i < 31; i++ {

			}
		}

	}
}
*/

//RETURN EVENTO WITH THE MOST PEOPLE
func EventoMostPeople() Evento {
	//VAR FOR THE EVENTO WITH THE MOST PEOPLE
	var a Evento
	//VALUE OF THE MOST PEOPLE IN EVENTOS SO FAR
	counttotal := 0
	//COUNT OF THE CURRENTE EVENTO
	countevento := 0
	//SEARCH IN THE ARRAY OF EVENTOS
	for i := 0; i < len(Lista_eventos); i++ {
		//SAVE THE VALUE OF NR PEOPLE IN EVENTO
		countevento = len(Lista_eventos[i].Bilhete)
		//IF HIGEST VALUE IS SMALLER THAN CURRENT VALUE REPLACE THE HIGEST VALUE
		if countevento > counttotal {
			counttotal = countevento
			//SAVE THE EVENTO WITH THE MOST PEOPLE
			a = Lista_eventos[i]
		}
	}
	//RETURN THE EVENTO
	return a
}

//COUNT OF HOW MUCH MONEY DID THEY WON
func CountAllMoney() float64 {
	//VAR TO COUNT THE MONEY
	var total float64
	//SEARCH THE ENTIRY ARRAY OF EVENTOS
	for i := 0; i < len(Lista_eventos); i++ {
		//MULTIPLY THE PRICEWITH THE NUMBER OF TICKETS
		total += Lista_eventos[i].Preço * float64(len(Lista_eventos[i].Bilhete))
	}
	//RETURN THE TOTAL
	return total
}

//COUNT OF HOW MUCH MONEY DID THEY WON IN CURRENT MONTH
func CountCurrentMonthMoney() float64 {
	//VAR TO COUNT THE MONEY
	var total float64
	time := time.Now()
	//SEARCH THE ENTIRY ARRAY OF EVENTOS
	for i := 0; i < len(Lista_eventos); i++ {
		//CHECK IF THE DATE IS THE SAME MONTH AS THE CURRENT AND YEAR
		if Lista_eventos[i].DataInicio.MesI == int(time.Month()) && Lista_eventos[i].DataFim.MesF == int(time.Month()) && Lista_eventos[i].DataInicio.AnoI == time.Year() && Lista_eventos[i].DataFim.AnoF == time.Year() {
			//MULTIPLY THE PRICEWITH THE NUMBER OF TICKETS
			total += Lista_eventos[i].Preço * float64(len(Lista_eventos[i].Bilhete))
		}
	}
	//RETURN THE TOTAL
	return total
}

//CHANGE THE PRICE OF THE TICKETS, IT CAN ONLY HAPPEN IF THERE ARE NO TICKET SOLD
func EditTicketPrice(nome string, price float64) int {
	//SEARCH THE ARRAY OF EVENTOS
	for i := 0; i < len(Lista_eventos); i++ {
		//CHECKS THE NAME OF THE EVENTO
		if nome == Lista_eventos[i].Nome {
			//CHECKS IF IT SOLD ANY TICKET
			if Lista_eventos[i].count == 0 {
				//CHANGE THE PRICE
				Lista_eventos[i].Preço = price
			} else {
				//RETURN 1 FOR FAILURE
				return 1
			}
		}
	}
	GuardarListaEventos()
	//RETURN 0 FOR SUCCESS
	return 0
}
