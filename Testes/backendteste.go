package main

import "fmt"

type ListaProfessores struct {
	lista_prof [100]Professor
}

type Get interface {
	getNota() float64
}

type Professor struct {
	Nome   string
	alunos [5]Aluno
}

type Aluno struct {
	Nome string
	ID   int
	nota float64
}

func (p Professor) getNota() float64 {
	for i := 0; i < len(p.alunos); i++ {
		return p.alunos[i].nota
	}
	return 1
}

func main() {

	var a1 Aluno
	a1.Nome = "Pedro"
	a1.ID = 919191
	a1.nota = 13.2

	var a2 Aluno
	a1.Nome = "Tiago"
	a1.ID = 121212
	a1.nota = 11.2

	var a3 Aluno
	a1.Nome = "Xico"
	a1.ID = 991
	a1.nota = 16.2

	var a4 Aluno
	a1.Nome = "Andre"
	a1.ID = 112
	a1.nota = 15.2

	var p1 Professor
	p1.Nome = "Alice"
	p1.alunos[0] = a1
	p1.alunos[1] = a2
	p1.alunos[2] = a3
	p1.alunos[3] = a4

	fmt.Println(p1.alunos[0].nota)
	fmt.Println(p1.alunos[1].nota)
	fmt.Println(p1.alunos[2].nota)
	fmt.Println(p1.alunos[3].nota)

	fmt.Println(p1.getNota())
}
