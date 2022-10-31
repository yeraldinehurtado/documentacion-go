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

func (e *estudiante) toPrint() { // metodo toPrint
	fmt.Printf("El estudiante %s \ntiene el siguiente horario: \n\t 1. %s \n\t 2. %s \n\t 3. %s", e.nombre, e.materia1, e.materia2, e.materia3)
}

func main4() {
	est1 := estudiante{
		persona: persona{
			nombre:   "Yeraldine",
			apellido: "Hurtado",
		},
		horario: horario{
			materia1: "Matemáticas",
			materia2: "Física",
			materia3: "Química",
		},
	}

	est1.toPrint()
}

/*
El estudiante Yeraldine 
tiene el siguiente horario: 
	 1. Matemáticas 
	 2. Física 
	 3. Química
*/