package main

import "fmt"

// función que se llama a sí misma

// con recursión
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// con ciclo
func cicloFac(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

func main() {
	//con recursion
	fmt.Println(factorial(7)) // factorial es como 7 * 6 * 5 * 4 * 3 * 2 * 1 = 5040

	// con ciclo
	fmt.Println(cicloFac(7))

}
