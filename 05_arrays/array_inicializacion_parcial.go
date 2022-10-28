package main

import "fmt"

func main6(){
	/* si se da un numero menor de elementos, entonces el
	compilador inicializa el resto al valor nulo.
	*/

	i := [5]int{5, 7, 9, 2}

	fmt.Println(i) // -> [5, 7, 9, 2, 0]
}