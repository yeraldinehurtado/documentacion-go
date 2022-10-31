package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x antes", x)
	fmt.Println("x antes", &x)
	foo(&x)
	fmt.Println("x después", x)
	fmt.Println("x después", &x)
}

func foo(y *int) {
	fmt.Println("y antes", y)
	fmt.Println("y antes", *y) // *y se llama de referencia y ahi estamos accediendo al valor almacenado en la dirección de memoria
	*y = 42
	fmt.Println("y después", y)
	fmt.Println("y después", *y)
}


// de esta manera podemos cambiar un valor en la dirección de memoria

/*
imprime ->

x antes 0
x antes 0xc00001c098
y antes 0xc00001c098
y antes 0
y después 0xc00001c098
y después 42
x después 42
x después 0xc00001c098

*/