package main

import (
	"fmt"
	"os"
)

func main2() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("El programa no finalizo de forma correcta")
		}
	}()

	if file, error := os.Open("Datos.txt"); error != nil {

		panic("Error: Archivo no encontrado")

	} else {

		contents := make([]byte, 1000)
		long, _ := file.Read(contents)
		contenidoarchivo := string(contents[:long])
		fmt.Println(contenidoarchivo)

		defer func() {
			fmt.Println("Cerrando archivos")
			file.Close()
		}()
	}

}
