package main

import (
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		// log.Println("un error ocurrió", err)
		// log.Fatalln(err)
		panic(err)
	}
}

/*
imprime ->

panic: open sin-archivo.txt: no such file or directory

goroutine 1 [running]:
main.main()
	/home/yeraldine/Documents/Mis proyectos GO/documentacion-go/15_errores/printing_logging/panic.go:13 +0x49
exit status 2
*/
