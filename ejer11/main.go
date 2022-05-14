package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}


type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
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
	vivo bool
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
func (h *hombre) estaVivo() bool { return h.vivo}



func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}



/*  GENERO ANIMAL   */

type perro struct {
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func (p *perro ) respirar()  { p.respirando = true }
func (p *perro ) comer()  { p.comiendo = true }
func (p *perro ) EsCarnivoro() bool  { return p.carnivoro }
func (p *perro ) estaVivo() bool  { return p.vivo }

func AnimalesRespirar(an animal)  {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/*  SER VIVO  */

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}


func main()  {

fmt.Println("------HUMANOS-----")

Pedro := new(hombre)
Pedro.esHombre = true
HumanoRespirando(Pedro)

Maria := new(mujer)
HumanoRespirando(Maria)


fmt.Println("------ANIMALES-----")



totalCarnivoros := 0

Dogo := new(perro)
Dogo.carnivoro = true
Dogo.vivo = true
AnimalesRespirar(Dogo)
totalCarnivoros =+ AnimalesCarnivoros(Dogo)


Pitbull := new(perro)
Pitbull.carnivoro = true
AnimalesRespirar(Pitbull)
totalCarnivoros += AnimalesCarnivoros(Pitbull)

fmt.Printf("Total Carnivoros %d \n", totalCarnivoros)

fmt.Printf("Estoy vivo = %t", estoyVivo(Dogo))

}