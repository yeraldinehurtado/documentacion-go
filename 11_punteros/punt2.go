package main

import "fmt"

func main() {
	var x *int
	y := 4

	x = &y // x almacena (o apunta a) la dirección de y

	// información de la variable x

	fmt.Printf("Lo que almacena la variable x -> %v \n", x)               // 0xc000018088
	fmt.Printf("Dirección en memoria de la variable x -> %v \n", &x)      // 0xc000012028
	fmt.Printf("Lo que almacena la dirección guardada en x -> %v \n", *x) // 4

	// información de la variable y

	fmt.Printf("Lo que almacena la variable y -> %v \n", y)          // 4
	fmt.Printf("Dirección en memoria de la variable y -> %v \n", &y) // 0xc000018088

}
