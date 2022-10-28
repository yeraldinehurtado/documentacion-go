package main

import (
	"fmt"
)

const s string = "constant"

func main() {

	/*
		Una constante es similar a una variables , pero no puede cambiar
		de valor. Se declaran de modo parecido a las variables pero utilizando
		la palabra reservada const
	*/

	fmt.Println(s) // imprime -> constant

	const pi = 3.141595

	fmt.Println(pi) // imprime -> 3.141595
}
