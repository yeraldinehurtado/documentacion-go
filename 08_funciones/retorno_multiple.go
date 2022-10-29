package main

import "fmt"

// La función devuelve la suma y la resta
func sumaResta(x int, y int) (int, int){
	return x + y, x-y
}

func main5(){
	a, b := sumaResta(67, 3)
	fmt.Println("La suma es", a, "y la resta es", b) // imprime -> La suma es 70 y la resta es 64
}