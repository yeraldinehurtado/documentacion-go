package main

import "fmt"

func main8() {
	
	y := [3]int{10, 20, 30}

	/* la palabra reservada range devuelve dos valores:
	- el primero es el índice
	- el segundo es el valor del dato */

	for indice, valor := range y {
		fmt.Println("El valor de la posición", indice, "es", valor)
	}

	// si queremos usar solo el indice o valor usamos como
	// variable el guion bajo

	for _, valor := range y { // no queremos el indice
		fmt.Println(valor)
	}
}