package main

import "fmt"

var tabla [10]int

var matriz [5][7]int

func main() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	miTabla := [10]int{5, 6, 98, 1, 4, 0, 3, 65, 8, 99}

	for i := 0; i < len(miTabla); i++ {
		fmt.Println(miTabla[i])
	}

	fmt.Println("--------------MATRIZ--------------------")

	matriz[3][5] = 1
	fmt.Println(matriz)

	fmt.Println("-----------MATRIZ DOS-----------------------")

	matrizDos := []int{2, 5, 4}
	fmt.Println(matrizDos)

	fmt.Println("--------------VARIANTE 2--------------------")
	variante2()

	fmt.Println("--------------VARIANTE 3--------------------")
	variante3()

	fmt.Println("--------------VARIANTE 4--------------------")
	variante4()

	fmt.Println("--------------VARIANTE 5--------------------")
	variente5()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4]
	fmt.Println(porcion)
}

func variante4() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

}

func variente5() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
