package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	licenciaPM bool
}

func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido)
}

func main4() {
	agenteSecreto1 := agenteSecreto{
		persona: persona{
			nombre:   "Yeraldine",
			apellido: "Hurtado",
		},
		licenciaPM: true,
	}

	fmt.Println(agenteSecreto1)
	agenteSecreto1.presentarse()
}
