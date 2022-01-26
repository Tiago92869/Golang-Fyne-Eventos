package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file2, ferr := os.Open("golang.txt")
	if ferr != nil {
		a := 5
		str := strconv.Itoa(a)

		file, err := os.Create("golang.txt")

		if err != nil {
			log.Fatal("We got error", err)
		}

		defer file.Close()

		fmt.Fprintf(file, str)
	}
	scanner := bufio.NewScanner(file2)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		fmt.Printf("O valor é: %s", items[0])
	}
}
