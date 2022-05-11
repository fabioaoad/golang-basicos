package main

import (
	"fmt"
)

var numero1 int
var numero2 int

func main() {
	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2: ")
	fmt.Scanln(&numero2)
	fmt.Printf("%d", numero1+numero2)
}
