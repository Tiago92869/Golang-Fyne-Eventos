package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var alunoNum, profNum = getNum()

//ALUNO
type Aluno struct {
	id   int
	nome string
	nota int
}

//PROFESSOR
type Professor struct {
	id     int
	nome   string
	alunos [2]Aluno
}

//LISTA DO ALUNO E COUNT
var (
	alunoLista []Aluno
	//alunoNum = 0
	//LISTA DO PROFESSOR E COUNT
	profLista []Professor
	//profNum = 0
	position = 0
)

//CRIAR UM NOVO ALUNO
func newAluno(id int, nome string, nota int) *Aluno {
	return &Aluno{
		id:   id,
		nome: nome,
		nota: nota,
	}
}

//CRIAR UM NOVO PROFESSOR
func newProf(id int, nome string) *Professor {
	return &Professor{
		id:   id,
		nome: nome,
	}
}

//DEVOLVER A LISTA DOS ALUNOS
func getAlunoList() []Aluno {
	return alunoLista
}

//DEVOLVER A LISTA DOS PROFESSORES
func getProfList() []Professor {
	return profLista
}

//verificar se existe pasta criada
func getNum() (int, int) {
	file2, ferr := os.Open("count.txt")

	if ferr != nil {
		alunoNum := 0
		profNum := 0
		str := strconv.Itoa(alunoNum)
		str2 := strconv.Itoa(profNum)

		file, err := os.Create("count.txt")

		if err != nil {
			log.Fatal("We got error", err)
		}

		defer file.Close()

		fmt.Fprintf(file, str)
		fmt.Fprint(file, ",")
		fmt.Fprintf(file, str2)
	}
	scanner := bufio.NewScanner(file2)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		s := items[0]
		ss := items[1]
		println(ss)
		println(s)
		i, _ := strconv.Atoi(s)
		ii, _ := strconv.Atoi(ss)

	}
	return i, ii
}

//ADICIONAR O ESTUDANTE AO
func adicionarEstudante() {
	var id = alunoNum
	alunoNum += 1 //soma mais 1 ao contador de alunos
	var nome string
	var nota int
	fmt.Print("Introduza o nome do aluno: ")
	_, err := fmt.Scan(&nome)
	fmt.Print("Introduza a nota do aluno: ")
	_, err = fmt.Scan(&nota)
	if err != nil {
		fmt.Println("Input error! ERROR: ", err)
	}
	newAluno(id, nome, nota)
	alunoLista = append(alunoLista, *newAluno(id, nome, nota))
}

//ADICIONAR O PROFESSOR À LISTA DE PROFESSORES
func adicionarProfessor() {
	var id = profNum
	profNum += 1 //soma mais 1 ao contador de alunos
	var nome string
	fmt.Print("Introduza o nome do professor: ")
	_, err := fmt.Scan(&nome)
	if err != nil {
		fmt.Println("Input error! ERROR: ", err)
	}
	newProf(id, nome)
	profLista = append(profLista, *newProf(id, nome))
}

//ELIMINAR UM NOVO PROFESSOR
func deleteAluno() { //Enter student id to delete
	var id int
	fmt.Println("Por favor introduza o id do aluno:")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println(" Input error! ERROR:", err)
	}
	for _, v := range alunoLista {
		if v.id == id {
			alunoLista = append(alunoLista[:id], alunoLista[id+1:]...) //append(value to replace, replace value...-> till the end)
			/*alunoLista[i] = alunoLista[len(a)-1]
			alunoLista[len(a)-1]= ""
			alunoLista = alunoLista[:len(a)-1]*/
		}
	}
}

//EDITAR UM ALUNO PELO ID
func editAluno() {
	var i int
	fmt.Println("Introduza o numero do id do estudante: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Input error! ERROR: ", err)
	}
	for _, v := range alunoLista {
		if v.id == i {
			fmt.Print("Introduza o nome do aluno: ")
			_, err := fmt.Scan(&alunoLista[i].nome)
			fmt.Print("Introduza a nota do aluno: ")
			_, err = fmt.Scan(&alunoLista[i].nota)
			if err != nil {
				fmt.Println("Input error! ERROR: ", err)
			}
		}
	}
}

//EDITAR UM ALUNO PELO ID
func editProf() {
	var i int
	fmt.Println("Introduza o numero do id do professor: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Input error! ERROR: ", err)
	}
	for _, v := range profLista {
		if v.id == i {
			fmt.Print("Introduza o nome do professor: ")
			_, err := fmt.Scan(&profLista[i].nome)
			if err != nil {
				fmt.Println("Input error! ERROR: ", err)
			}
		}
	}
}

func adicionarAlunoParaProfessor() {
	var ii int
	fmt.Println("Introduza o numero do id do professor: ")
	_, err := fmt.Scan(&ii)
	if err != nil {
		fmt.Println("Input error! ERROR: ", err)
	}
	for _, vv := range profLista {
		if vv.id == ii {
			var i int
			fmt.Println("Introduza o numero do id do estudante: ")
			_, err = fmt.Scan(&i)
			if err != nil {
				fmt.Println("Input error! ERROR: ", err)
			}
			for _, v := range alunoLista {
				if v.id == i {
					profLista[ii].alunos[position] = alunoLista[i]
					position++
				}
			}
		}
	}
}

func main() {
	/*
		var a1 Aluno
		a1.nome = "Constantina"
		a1.nota = 12

		var a2 Aluno
		a2.nome = "Elsa"
		a2.nota = 20

		var prof Professor
		prof.nome = "Helena"
		prof.alunos[0] = a1
		prof.alunos[1] = a2

		println(prof.alunos[0].nome)
		println(prof.alunos[0].nota)
		println(prof.alunos[1].nome)
		println(prof.alunos[1].nota)
	*/
	//FC CODERS AZUREM
	//codigo para escrever em txt
	println("aluno", alunoNum)
	println("profe", profNum)
	println(alunoNum)
	adicionarProfessor()
	fmt.Println(getProfList())
	adicionarEstudante()
	fmt.Println(getAlunoList())
	adicionarAlunoParaProfessor()
	println(profLista[0].alunos[0].nome)
	println(profLista[0].alunos[0].nota)

	f, err := os.Create("alunos.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, v := range alunoLista {
		var aluno string
		aluno = v.nome + "\n"
		f.WriteString(aluno)
		aluno = strconv.Itoa(v.id) + "\n"
		f.WriteString(aluno)
		aluno = strconv.Itoa(v.nota) + "\n"
		f.WriteString(aluno)
	}
}
