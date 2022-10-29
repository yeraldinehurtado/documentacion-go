package main

import "fmt"

// La función recibe dos int y devuelve un int
func area(base int, altura int) int {
	resultado := base * altura
	return resultado
}

func main() {
	a := area(45, 90)
	fmt.Println(a)                              // imprime -> 4050
	fmt.Println("Ahora el área es", area(7, 9)) // imprime -> Ahora el área es 63
	fmt.Printf("El tipo es: %T\n", area)        // imprime -> El tipo es: func(int, int) int
}
