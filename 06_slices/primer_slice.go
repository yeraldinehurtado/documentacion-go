package main

import "fmt"

func main2() {

	z := [...]int{1, 2, 3, 4, 5} // imprime -> [1 2 3 4 5]
	z1 := z[:]                   // imprime -> [1 2 3 4 5]
	z2 := z[2:]                  // imprime -> [3 4 5]
	z3 := z[:3]                  // imprime -> [1 2 3]
	z4 := z[0:4]                 // imprime -> [1 2 3 4]
	z5 := z[1:4]                 // imprime -> [2 3 4]

	fmt.Println(z)
	fmt.Println(z1)
	fmt.Println(z2)
	fmt.Println(z3)
	fmt.Println(z4)
	fmt.Println(z5)

	fmt.Println(" \nAntes de usar append()\n ")
	fmt.Printf("slice z -> %v -> len %v || cap -> %v \n", z, len(z), cap(z))
	// imprime -> slice z -> [1 2 3 4 5] -> len 5 || cap -> 5

	fmt.Printf("slice z1 -> %v -> len %v || cap -> %v \n", z1, len(z1), cap(z1))
	// imprime -> slice z1 -> [1 2 3 4 5] -> len 5 || cap -> 5

	fmt.Printf("slice z2 -> %v -> len %v || cap -> %v \n", z2, len(z2), cap(z2))
	// imprime -> slice z2 -> [3 4 5] -> len 3 || cap -> 3

	fmt.Printf("slice z3 -> %v -> len %v || cap -> %v \n", z3, len(z3), cap(z3))
	// imprime -> slice z3 -> [1 2 3] -> len 3 || cap -> 5

	fmt.Printf("slice z4 -> %v -> len %v || cap -> %v \n", z4, len(z4), cap(z4))
	// imprime -> slice z4 -> [1 2 3 4] -> len 4 || cap -> 5

	fmt.Printf("slice z5 -> %v -> len %v || cap -> %v \n", z5, len(z5), cap(z5))
	// imprime -> slice z5 -> [2 3 4] -> len 3 || cap -> 4

	fmt.Println(" \nUsando append()\n ")

	z1 = append(z1, 6, 7, 8, 9, 10)
	z2 = append(z2, 6, 7, 8, 9, 10, 11, 12, 13)
	z3 = append(z3, []int{3, 4, 5, 6, 7, 8, 9}...)
	z4 = append(z4, []int{5, 6, 7}...)
	z5 = append(z5, z2...)

	fmt.Println(" \nDespués de usar append()\n ")

	fmt.Printf("slice z -> %v -> len %v || cap -> %v \n", z, len(z), cap(z))
	// imprime -> slice z -> [1 2 3 4 5] -> len 5 || cap -> 5

	fmt.Printf("slice z1 -> %v -> len %v || cap -> %v \n", z1, len(z1), cap(z1))
	// imprime -> slice z1 -> [1 2 3 4 5 6 7 8 9 10] -> len 10 || cap -> 10

	fmt.Printf("slice z2 -> %v -> len %v || cap -> %v \n", z2, len(z2), cap(z2))
	// imprime -> slice z2 -> [3 4 5 6 7 8 9 10 11 12 13] -> len 11 || cap -> 12

	fmt.Printf("slice z3 -> %v -> len %v || cap -> %v \n", z3, len(z3), cap(z3))
	// imprime -> slice z3 -> [1 2 3 3 4 5 6 7 8 9] -> len 10 || cap -> 10

	fmt.Printf("slice z4 -> %v -> len %v || cap -> %v \n", z4, len(z4), cap(z4))
	// imprime -> slice z4 -> [1 2 3 4 5 6 7] -> len 7 || cap -> 10

	fmt.Printf("slice z5 -> %v -> len %v || cap -> %v \n", z5, len(z5), cap(z5))
	// imprime -> slice z5 -> [2 3 4 3 4 5 6 7 8 9 10 11 12 13] -> len 14 || cap -> 14


}
