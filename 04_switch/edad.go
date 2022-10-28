package main

import "fmt"

func main2() {
	// este programa tiene cuatro casos para comprobar y un caso por defecto.

	var edad = 14 // imprime -> adolescente

	switch {
	case edad < 0:
		fmt.Println("No ha nacido")
	case edad < 12:
		fmt.Println("Niño")
	case edad < 18:
		fmt.Println("Adolescente")
	case edad < 65:
		fmt.Println("Adulto")
	default:
		fmt.Println("Tercera edad")
	}
}
