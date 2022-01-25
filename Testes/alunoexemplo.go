package main

import (
	"fmt"
)

type Student struct {
	id    int
	name  string
	age   int8
	score float32
}

var (
	stuList []Student
	studNum = 0 //used to generate id
)

func newStu(id int, name string, age int8, score float32) *Student { //Construction method
	return &Student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

func getStuList() []Student { //Return a list of students, used to query
	return stuList
}

func addStu() { //add student
	var id = studNum
	studNum += 1 //id increases with the number of students
	var name string
	var age int8
	var score float32
	fmt.Print("Please enter student name:")
	_, err := fmt.Scan(&name)
	fmt.Print(" Please enter the age of the student: ")
	_, err = fmt.Scan(&age)
	fmt.Print("Please enter student score:")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("Input error! ERROR:", err)
	}
	newStu(id, name, age, score)
	stuList = append(stuList, *newStu(id, name, age, score))
}

func deleteStu() { //Enter student id to delete
	var i int
	fmt.Println("Please enter student id:")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println(" Input error! ERROR:", err)
	}
	for _, v := range stuList {
		if v.id == i {
			stuList = append(stuList[:i], stuList[i+1:]...)
		}
	}
}

func editInfo() { //Enter the student id to edit and modify it
	var i int
	fmt.Println("Please enter the student id:")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("Input error! ERROR:", err)
	}
	for _, v := range stuList {
		if v.id == i {
			fmt.Print("Please enter the student's name:")
			_, err := fmt.Scan(&stuList[i].name)
			fmt.Print("Please enter the student's age:")
			_, err = fmt.Scan(&stuList[i].age)
			fmt.Print("Please enter student grades:")
			_, err = fmt.Scan(&stuList[i].score)
			if err != nil {
				fmt.Println("Input error! ERROR:", err)
			}
		}
	}
}

func main() {
	for {
		fmt.Println("The operation to be performed:")
		fmt.Print("1. Add 2. View 3. Delete 4. Modify\n")
		var do int8
		_, err := fmt.Scan(&do)
		if err != nil {
			fmt.Println("The input is wrong!")
		}
		switch do {
		case 1:
			addStu()
		case 2:
			fmt.Println(getStuList())
		case 3:
			deleteStu()
		case 4:
			editInfo()
		default:
			fmt.Println("input error!")
		}
	}
}
