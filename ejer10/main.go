package main

import (
	"fmt"
	us "golang-basicos/ejer10/user"
	"time"
)

type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

type martin struct {
	us.Usuario
}

func main() {
	User := new(usuario)
	User.Id = 10
	User.Nombre = "Fabio"
	fmt.Println(User)

	fmt.Println("--------------------")

	u := new(martin)
	u.AltaUsuario(1, "Fabio Aoad", time.Now(), true)
	fmt.Println(u)

}
