package main

//CONSTRUTOR DA DATA
type Data struct {
	Ano    int
	Mes    int
	Dia    int
	Hora   int
	Minuto int
}

//CONSTRUTOR DOS BILHETES
type Bilhete struct {
	ID int
}

//CONSTRUTOR DOS EVENTOS
type Evento struct {
	Nome          string
	Data          Data
	Duração       int
	Participantes int
	Bilhete       []Bilhete
	Preço         float64
}

//CONSTRUTOR DOS PARTICIPANTES
type Participante struct {
	Nome    string
	Bilhete Bilhete
}

var lista_eventos []Evento

func registoEvento(nome string, ano int, mes int, dia int, hora int, minutos int, duração int, participantes int, preço float64) *Evento {
	evento := Evento{Nome: nome}
	evento.Data.Ano = ano
	evento.Data.Mes = mes
	evento.Data.Dia = dia
	evento.Data.Hora = hora
	evento.Data.Minuto = minutos
	evento.Duração = duração
	evento.Participantes = participantes
	evento.Preço = preço
	return &evento
}

func getEvento(nome string, lista lista_eventos) *Evento {

}

func main() {

}
