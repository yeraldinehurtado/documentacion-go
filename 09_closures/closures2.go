package main

import "fmt"

func numeros() func() int {
	i := 0                   
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := numeros()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := numeros()
	fmt.Println(newInts())
}
