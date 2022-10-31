package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func main2() {
	fmt.Println(persona{"Bob", 20}) // Imprime -> {"Bob", 20}

	fmt.Println(persona{nombre: "Alice", edad: 30}) // Imprime -> {Alice 30}

	fmt.Println(persona{nombre: "Fred"}) // Imprime -> {Fred 0}

	fmt.Println(&persona{nombre: "Ann", edad: 40}) // Imprime &{Ann 40}
	// El prefijo & devuelve el apuntador a un struct.

	s := persona{nombre: "Sean", edad: 50} 
	fmt.Println(s.nombre) // Imprime -> Sean

	sp := &s
	fmt.Println(sp.edad) // Imprime -> 50

	sp.edad = 51
	fmt.Println(sp.edad) // Imprime -> 51
}
