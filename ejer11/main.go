package main

import "fmt"

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}


type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal()
}


/*  GENERO HUMANO   */

type hombre struct {
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true}
func (h *hombre) comer() { h.comiendo = true}
func (h *hombre) pensar() { h.pensando = true}
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}



func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

func main()  {

Pedro := new(hombre)
Pedro.esHombre = true
HumanoRespirando(Pedro)

Maria := new(mujer)
HumanoRespirando(Maria)


}