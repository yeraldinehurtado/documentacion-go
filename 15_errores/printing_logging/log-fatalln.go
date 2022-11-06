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
		log.Fatalln(err)
		// panic(err)
	}
}

// imprime -> 2022/11/06 13:46:21 open sin-archivo.txt: no such file or directory
// exit status 1
