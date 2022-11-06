package main

import (
	"fmt"
	"os"
)

func main(){
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		fmt.Println("err happened", err)
		// log.Println("un error ocurrió", err)
		// log.Fatalln(err)
		// panic(err)
	}
}

// imprime -> err happened open sin-archivo.txt: no such file or directory