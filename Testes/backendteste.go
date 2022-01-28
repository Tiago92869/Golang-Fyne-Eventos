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
	id       int
	nome     string
	alunos   [2]Aluno
	position int
}

//LISTA DO ALUNO E COUNT
var alunoLista []Aluno = carregarListaAlunos()
var profLista []Professor

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
		id:       id,
		nome:     nome,
		position: 0,
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

//CRIAR A PASTA DE ALUNOS SE NÃO EXISTIR
func CriarListaAlunos() {
	_, ferr := os.Open("alunos.txt")
	if ferr != nil {
		file, err := os.Create("alunos.txt")
		if err != nil {
			log.Fatal("We got error", err)
		}
		defer file.Close()
	}
}

//Criar a pasta alunos.txt se esta não existir, mas se existir guarda os dados dos alunos
func GuardarListaAlunos() {
	file2, _ := os.Create("alunos.txt")
	size := len(alunoLista)
	for i := 0; i < size; i++ {
		var aluno string
		aluno = strconv.Itoa(alunoLista[i].id) + ","
		file2.WriteString(aluno)
		aluno = alunoLista[i].nome + ","
		file2.WriteString(aluno)
		aluno = strconv.Itoa(alunoLista[i].nota) + "\n"
		file2.WriteString(aluno)
	}
}

//carregar a lista de alunos guardada para a lista de aluno
func carregarListaAlunos() []Aluno {
	var list []Aluno
	file, err := os.Open("alunos.txt")
	if err != nil {
		log.Fatal("We got error", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		//scanner por linha
		line := fileScanner.Text()
		//separar os elementos por virgula
		items := strings.Split(line, ",")
		//guarda os valores do txt em variáveis string
		s, _ := strconv.Atoi(items[0])
		println(s)
		ss := items[1]
		println(ss)
		sss, _ := strconv.Atoi(items[2])
		list = append(list, *newAluno(s, ss, sss))
	}
	return list
}

//verificar se existe pasta criada
//se nao existir criamos uma nova pasta onde iremos inserir o numero de alunos e o numero de professores e a posição respectivamente
//se existir iremos returnar o numero de alunos e o numero de professores
func getNum() (int, int) {
	//file2 corresponde à ação de abrir o ficheiro txt
	file2, ferr := os.Open("count.txt")
	//criamos os pointers para os valores da pasta
	var i *int
	var ii *int
	//se não existir o ficheiro, criamos um novo
	if ferr != nil {
		//convertemos os valores int para string
		str := strconv.Itoa(0)
		str2 := strconv.Itoa(0)
		//criação da nova pasta
		file, err := os.Create("count.txt")
		//verificar se não ocorreu nenhum erro na criação
		if err != nil {
			log.Fatal("We got error", err)
		}

		defer file.Close()
		//inserimos os ficheiros dentro da pasta
		fmt.Fprintf(file, str)
		fmt.Fprint(file, ",")
		fmt.Fprintf(file, str2)
	}
	scanner := bufio.NewScanner(file2)
	//lê-mos os valores guardados no ficheiro
	for scanner.Scan() {
		//scanner por linha
		line := scanner.Text()
		//separar os elementos por virgula
		items := strings.Split(line, ",")
		//guarda os valores do txt em variáveis string
		s := items[0]
		ss := items[1]
		//converte os valores para int
		si, _ := strconv.Atoi(s)
		ssi, _ := strconv.Atoi(ss)
		//guarda os valores no spointers
		i = &si
		ii = &ssi
	}
	return *i, *ii
}

//ADICIONAR O ESTUDANTE
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
	GuardarListaAlunos()
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
	GuardarListaAlunos()
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
	GuardarListaAlunos()
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

//adicionar um aluno á lista de alunos de um respetivo professor
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
					profLista[ii].alunos[profLista[ii].position] = alunoLista[i]
					profLista[ii].position = profLista[ii].position + 1
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
	CriarListaAlunos()
	println("aluno", alunoNum)
	println("profe", profNum)
	println(alunoNum)
	//adicionarEstudante()
	//alunoLista = carregarListaAlunos()
	fmt.Println(alunoLista[0])
	fmt.Println(alunoLista[1])
	fmt.Println(alunoLista[2])
	fmt.Println(alunoLista[5])

	println(len(alunoLista))
	adicionarEstudante()
	println(len(alunoLista))
	adicionarEstudante()
	println(len(alunoLista))
	adicionarEstudante()
	println(len(alunoLista))
	fmt.Println(getAlunoList())

	//adicionarAlunoParaProfessor()
}
