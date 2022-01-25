package main

import "fmt"

//lista professores
type ListaProfessores struct {
	lista_prof [100]Professor
}

//interface das funções
type Get interface {
	getNota() float64
	newAluno() Aluno
	newProfessor() Professor
}

//estrutura professor
type Professor struct {
	Nome   string
	alunos [100]Aluno
}

//estrutura aluno
type Aluno struct {
	Nome string
	ID   int
	nota float64
}

//lista de todos alunos
var (
	stuList []Aluno
	studNum = 0
)

//novo aluno
func newAluno(nome string, ID int, nota float64) *Aluno {
	return &Aluno{
		Nome: nome,
		ID:   ID,
		nota: nota,
	}
}

//RETURN NOVO PROFESSOR
func newProfessor(nome string) *Professor {
	return &Professor{
		Nome: nome,
	}
}

//RETURN DA NOTA DOS ALUNOS DE UM PROFESSOR
//AINDA NAO FUNCIONA
func (p Professor) getNota() float64 {
	for i := 0; i < len(p.alunos); i++ {
		return p.alunos[i].nota
	}
	return 1
}

func main() {

	var a1 = newAluno("Pedro", 919191, 13.2)

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
	p1.alunos[1] = a2
	p1.alunos[2] = a3
	p1.alunos[3] = a4

	fmt.Println(p1.alunos[0].nota)
	fmt.Println(p1.alunos[1].nota)
	fmt.Println(p1.alunos[2].nota)
	fmt.Println(p1.alunos[3].nota)

	//FUNÇÃO DE ADICIONAR ALUNO A FUNCIONAR
	var a5 = newAluno("Joao", 1234567, 12)
	fmt.Println(a5.nota)
	fmt.Println(a5.ID)
	fmt.Println(a5.Nome)

	fmt.Println(p1.getNota())
}
