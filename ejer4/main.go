package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2: ")
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion: ")

	scaner := bufio.NewScanner(os.Stdin) // escanea por teclado, es deir estandar input
	if scaner.Scan() {
		leyenda = scaner.Text()
	}

	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
