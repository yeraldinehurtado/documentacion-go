package main

import "fmt"

func main3() {

	// funcion len(): sirve para recorrer el número de elementos de un array

	u := [5]int{10, 20, 30, 40, 50}
	fmt.Println(len(u)) // -> 5

	//recorrer los elementos de un array con un bucle for

	y := [3]int{10, 20, 30}
	for i := 0; i < len(y); i += 1 {
		fmt.Println("El valor de la posición ", i, "es", y[i])
	}

}