package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
)

func main()  {
	leoArchivo()

	fmt.Println("-----------------------")
	leoArchivo2()

	fmt.Println("-----------------------")
	graboArchivo()


	fmt.Println("-----------------------")
	graboArchivo2()
}

func leoArchivo()  {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil{
		fmt.Println("Hubo un error")
	}else{
		fmt.Println(string(archivo))
	}
}


func leoArchivo2()  {
	archivo, err := os.Open("./archivo.txt")
	if err != nil{
		fmt.Println("Hubo un error")
	}else{
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro:= scanner.Text()
			fmt.Printf("Linea > "+registro+"\n")
		}
	}
	archivo.Close()
}

func graboArchivo()  {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil{
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

func graboArchivo2()  {
	fileName:= "./nuevoArchivo.txt"
	if Append(fileName, "\n Esta es una segunda linea") == false {
		fmt.Println("Error en la 2da linea")
	}
}

func Append(archivo string, texto string) bool  {
	arch, err := os.OpenFile(archivo, os.O_WRONLY| os.O_APPEND, 0644)
	if err != nil{
		fmt.Println("Hubo un error")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil{
		fmt.Println("Hubo un error")
		return false
	}
	return true
}