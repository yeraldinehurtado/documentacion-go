package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type horario struct {
	materia1 string
	materia2 string
	materia3 string
}

type estudiante struct {
	persona
	horario
}

func (e estudiante) toPrint() { // metodo toPrint
	fmt.Printf("El estudiante %s \ntiene el siguiente horario: \n\t 1. %s \n\t 2. %s \n\t 3. %s", e.nombre, e.materia1, e.materia2, e.materia3)
}

func (p persona) toPrint() { // metodo toPrint
	fmt.Printf("Saluda la persona")
}

// estudiante es de tipo humano
type humano interface { // colección de métodos
	toPrint()
}

func bar(h humano) {
	fmt.Println("Fui pasado a la función bar", h)
}

func main() {
	est1 := estudiante{
		persona: persona{
			nombre:   "Yeraldine",
			apellido: "Hurtado",
		},
		horario: horario{
			materia1: "Matemáticas",
			materia2: "Física",
			materia3: "Química\n",
		},
	}

	p := persona {
		nombre: "Elizabeth",
		apellido: "Muñoz",
	}

	est1.toPrint()

	bar(est1)
	bar(p)
}

// imprime ->

/*
El estudiante Yeraldine 
tiene el siguiente horario: 
	 1. Matemáticas 
	 2. Física 
	 3. Química
Fui pasado a la función bar {{Yeraldine Hurtado} {Matemáticas Física Química
}}
Fui pasado a la función bar {Elizabeth Muñoz}
*/