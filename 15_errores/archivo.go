package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("nombres.txt") // create retorna un puntero a un archivo y un error
	if err != nil { // manejo de error 
		fmt.Println(err)
		return
	}
	defer f.Close() // cada que se abre un archivo, hay que cerrarlo al final del programa para que no se quede
	// consumiendo recursos necesarios

	r := strings.NewReader("yeraldine") // newreader me retorna un puntero a un reader y le pasamos un string "yeraldine"

	io.Copy(f, r) // luego con la funcion copy estoy escribiendo el reader de arriba en el archivo 
}
