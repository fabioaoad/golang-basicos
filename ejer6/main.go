package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	fmt.Println(tres(5))

	numero, estado := dos(2)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println("----CALCULO------")

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(2, 23, 45, 68))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5, 46, 17, 25, 26, 98, 17))

	fmt.Println("-----CALCULO OTRO-----")

	fmt.Println(calculoOtro(5, 46))
	fmt.Println(calculoOtro(2, 23, 45, 68))
	fmt.Println(calculoOtro(5))
	fmt.Println(calculoOtro(5, 46, 17, 25, 26, 98, 17))

}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func tres(numero int) (z int) {
	z = numero * 2
	return
}

// Parámetros variables, util ya que go no existe sobrecarga de metodos en go

func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		// range retorna 2 valores, el primero es el numero del elemento, es decir,
		// si es el uno, dos, etc. Lo almaceno en _ pero no lo usaré.
		// El segundo es
		total = total + num
	}
	return total
}

func calculoOtro(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
