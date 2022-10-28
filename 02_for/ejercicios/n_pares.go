package main

import "fmt"

func main2() {
	// Crea un programa que imprima los números pares del 10 al 50

	for i := 10; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			i++
		}
	}
	fmt.Println("Programa terminado")
}
