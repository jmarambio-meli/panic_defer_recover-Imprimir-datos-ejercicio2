package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	defer func() {
		fmt.Println("ejecución finalizada")
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Open("customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		fmt.Printf("Legajo: %s, Name: %s, DNI: %s, Telefono: %s, Direccion: %s \n", items[0], items[1], items[2], items[3], items[4])
	}
}
