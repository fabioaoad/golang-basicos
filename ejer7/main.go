package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n ", Calculo(5, 7))

	// restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resto 6 - 4 = %d \n ", Calculo(6, 4))

	// divido
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Divido 12 / 3 = %d \n ", Calculo(12, 3))

	// multiplico
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Printf("Multiplico 5 * 6 = %d \n ", Calculo(5, 6))

	//funcion anonima
	Operaciones()

	// closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

//funciones anonimas

//creo una funcion la cual contiene una variable de tipo
//funcion que retorna un valor

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

// closures

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia = secuencia + 1
		return numero * secuencia
	}
}
