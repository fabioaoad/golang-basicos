package main

import "fmt"

func main() {

	// for basico
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// for con break
	numero := 0
	for {
		fmt.Println("continuo")
		fmt.Println("ingrese el númer secreto: ")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf(" Multiplicamos por 2 \n")
			i = i * 2
			continue // vuelvo a la condicion del for
		}
		fmt.Printf("    pasó por aquí \n")
		i++
	}

	var j int = 0

RUTINA:
	for j < 10 {
		if j == 4 {
			j = j + 2
			fmt.Println("voy a RUTINA sumando 2 a j")
			goto RUTINA // similar a continuo con la diferencia que irá a donde yo le diga, es decir, donde tenga la etiqueta RUTINA
		}
		fmt.Printf("Valor de j: %d\n", j)
		j++
	}

}
