package main

import (
	"fmt"
	"strconv"
)

//var numero int
//var texto string
var status bool = true // se inicializa en false si no le indico

func main() {
	fmt.Println("Hola Mundo")
	var num1, num2, num3 int
	var moneda int64 = 323
	var moneda2 float64 = 0
	num1 = int(moneda)
	num1, num2, num3, texto, texto2, status := 34, 54, 66, "este es mi texto", "mi otro texto2", false
	texto = fmt.Sprintf("%d", moneda) // muestra como string a 323
	texto2 = strconv.Itoa(int(moneda2))
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(texto)
	fmt.Println(status)
	mostrarStatus()
	fmt.Println(moneda)
	fmt.Println(texto2)

}

func mostrarStatus() {
	fmt.Println(status)
}
