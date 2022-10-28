package main

import "fmt"

func main() {
	// Comprobar si un número entero es positivo, negativo o nulo

	x := 2 // El resultado de la condición es el número es positivo

	if x > 0 {
		fmt.Println("El número es positivo")
	}
	if x == 0 {
		fmt.Println("El número es nulo")
	}
	if x < 0 {
		fmt.Println("El número es negativo")
	}
}
