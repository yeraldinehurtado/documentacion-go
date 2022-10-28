package main

import "fmt"

func main2(){

	u := [5]int{24, 87, 72, 34, 21}     //tipo int y 5 datos
	v := [...]float64{10.1, 38.3, 27.3} //número de datos lo puede inferir el compilador con ...

	fmt.Printf("%v es de tipo %T\n", u, u)
	fmt.Printf("%v es de tipo %T\n", v, v)

	//acceder a los datos de manera individual

	fmt.Println(u[0]) // -> 24
	fmt.Println(u[1]) // -> 87
	fmt.Println(u[4]) // -> 21

	// La misma anotación se puede hacer para cambiar el valor de algún vector

	u[0] = 999
	fmt.Println(u) // -> [999, 87, 72, 34, 21]

	// Para acceder a varios elementos del vector se usa :

	fmt.Println(u[0:3]) // -> [999, 87, 72]
	fmt.Println(u[1:4]) // -> [87, 72, 34]




	// ----
	// ----
	// ----
	// ----

	fmt.Println("\nEsto es de http://goconejemplos.com/arreglos\n ")

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}