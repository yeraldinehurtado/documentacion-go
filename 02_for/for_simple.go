package main

import "fmt"

func main2() {

	// Hola no se ejecuta ninguna vez
	for false {
		fmt.Println("Hola")
	}

	// Bucle infinito de Adiós
	/*	for true {
			fmt.Println("Adiós")
		}
	*/

	// Imprime los número del 1 al 10
	i := 1 //inicializa variable

	for i <= 10 {
		fmt.Println(i)
		i++          // alguna vez i será 11
	}
	fmt.Println("Programa terminado")

}
