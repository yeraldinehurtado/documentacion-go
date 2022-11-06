package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		// log.Println("un error ocurrió", err)
		// log.Fatalln(err)
		log.Panicln(err)
		//panic(err)
	}
}

/*
imprime ->

2022/11/06 15:02:15 open sin-archivo.txt: no such file or directory
panic: open sin-archivo.txt: no such file or directory


goroutine 1 [running]:
log.Panicln({0xc00009af60?, 0xf?, 0x0?})
	/usr/local/go/src/log/log.go:402 +0x65
main.main()
	/home/yeraldine/Documents/Mis proyectos GO/documentacion-go/15_errores/printing_logging/panicln.go:14 +0x56
exit status 2
*/
