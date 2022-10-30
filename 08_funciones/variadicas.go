package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ") // para mostrar los numeros que se suman
	total := 0
	for _, num := range nums {
		total += num // suma los numeros ingresados
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)    // imprime -> [1 2] 3
	sum(1, 2, 3) // imprime -> [1 2 3] 6

	nums := []int{1, 2, 3, 4}
	sum(nums...) // imprime -> [1 2 3 4] 10
}
