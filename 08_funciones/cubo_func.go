package main

import "fmt"

func cubo(numero int) {
	var potencia = numero * numero * numero
	fmt.Println("El cubo de", numero, "es", potencia)
}

func main4() {
	cubo(2) // imprime -> el cubo de 2 es 8
}
