package main

import "fmt"

func cero(ival int) { // NO puntero
	ival = 0
}

func ceroPuntero(ipuntero *int) { // Recibe un apuntador a un valor int
	*ipuntero = 0 // Puntero
}

func main() {
	i := 1
	fmt.Println("Inicial: ", i) // imprime -> Inicial: 1

	cero(i)
	fmt.Println(i) // NO puntero, imprime -> 1

	ceroPuntero(&i) // paso i = 1 
	fmt.Println(i) // imprime -> 0

	fmt.Println("Puntero: ", &i) // ubicación puntero en memoria
	// imprime -> Puntero:  0xc0000b4000
}
