package main

import "fmt"

func main3() {
	//declaramos variables
	hola := "Hola"
	mundo := "Mundo"

	//imprimimos
	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "yeraldine"
	edad := 19
	fmt.Printf("Hola, %s y su edad %d\n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %s y su edad %d", nombre, edad)

	fmt.Println(mensaje)

	//ingresar datos en consola
	fmt.Println("ingrese otro nombre")
	fmt.Scan(&nombre)

	fmt.Println("Otro nombre:", nombre)
}
