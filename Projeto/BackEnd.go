package main

type Data struct {
	Ano    int
	Mes    int
	Dia    int
	Hora   int
	Minuto int
}

type Bilhete struct {
	ID int
}

type Evento struct {
	Nome string
	Data
	Duração       int
	Participantes int
	Bilhete
	Preço float64
}

type Participante struct {
	Nome    string
	Bilhete Bilhete
}

func registoEvento(nome string, ano int, mes int, dia int, hora int, minutos int, duração int, participantes int, Bilhete, preço float64) *Evento {
	evento := Evento{Nome: nome}
	evento.Ano = ano
	evento.Mes = mes
	evento.Dia = dia
	evento.Hora = hora
	evento.Minuto = minutos
	evento.Duração = duração
	evento.Participantes = participantes
	evento.Bilhete = Bilhete
	evento.Preço = preço

}
