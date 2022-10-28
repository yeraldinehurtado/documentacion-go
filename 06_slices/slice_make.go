package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("slice vacío:", s)

	s[0] = "a" // ingresando elementos al slice s
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Contenido slice:", s)
	fmt.Println("Posicion 2 del slice:", s[2])

	fmt.Println("Cantidad de elementos:", len(s))

	s = append(s, "d") // agregando mas elementos con append
	s = append(s, "e", "f")

	fmt.Println("Slice con nuevos elementos:", s)

	c := make([]string, len(s)) // aqui creamos un slice c vacio del mismo tamaño que s
	copy(c, s)                  // copiamos el contenido de s en c
	fmt.Println("slice c con elementos del slice s:", c)
}
