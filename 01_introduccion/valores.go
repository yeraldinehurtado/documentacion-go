package main

import "fmt"

func main2() {

	/* go maneja varios tipos de valores incluyendo cadenas,
	enteros, flotantes, booleanos, etc. */

	//cadenas de texto concatenadas con +
	fmt.Println("Go" + "lang")

	//enteros y flotantes (int y float)
	fmt.Println("2 + 2 = ", 2+2)
	fmt.Println("5.0/3.0 = ", 5.0/3.0)

	//booleanos, con operadores booleanos.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
