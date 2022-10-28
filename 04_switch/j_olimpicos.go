package main

import "fmt"

func main() {
	// se pueden poner varios resultados en un mismo caso:

	posicion := 7 // imprime diploma olimpico

	switch posicion {
	case 1:
		fmt.Println("Oro")
	case 2:
		fmt.Println("Plata")
	case 3:
		fmt.Println("Bronce")
	case 4, 5, 6, 7, 8:
		fmt.Println("Diploma olimpico")
	default:
		fmt.Println("Lo importante es participar :)")
	}
}
