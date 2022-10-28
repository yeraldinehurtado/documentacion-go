package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Euler":   1707,
		"Riemann": 1826,
	}
	fmt.Println(m) // imprime -> map[Euler:1707 Riemann:1826]
	fmt.Printf("%T\n", m) // imprime -> map[string]int

	//diccionarios
	nombre := make(map[int][]string) // mapa multidimensional
	nombre[0] = []string{
		"yeraldine", "hurtado",
		"juan", "cochero",
	}

	for i := 0; i < len(nombre); i++ { // recorriendo mapa con for
		fmt.Println("Nombres: ", nombre[i]) // imprime -> Nombres:  [yeraldine hurtado juan cochero]
	}

	for clave, valor := range nombre { //recorriendo con for range
		fmt.Println("Clave: ", clave, "valor: ", valor) // imprime -> Clave:  0 valor:  [yeraldine hurtado juan cochero]
	}
}
