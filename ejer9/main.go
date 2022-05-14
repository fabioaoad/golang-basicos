package main

import "fmt"

func main() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises["Mexico"])

	fmt.Println("-----------------")

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	fmt.Println(campeonato)

	fmt.Println("-----------------")

	campeonato["River Plate"] = 25

	fmt.Println(campeonato)

	fmt.Println("-----------------")

	campeonato["Chivas"] = 55
	fmt.Println(campeonato)

	fmt.Println("-----------------")

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	fmt.Println("-----------------")

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de : %d \n", equipo, puntaje)
	}

	fmt.Println("-----------------")

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capurado es %d, y el equipo existe %t \n", puntaje, existe)

	fmt.Println("-----------------")

	puntaje2, existe2 := campeonato["Chivas"]
	fmt.Printf("El puntaje capurado es %d, y el equipo existe %t \n", puntaje2, existe2)
}
