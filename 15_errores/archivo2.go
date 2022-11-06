package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	f, err := os.Open("nombres.txt") // con open abrimos el archivo
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // esto es para que se cierre el archivo al final

	bs, err := ioutil.ReadAll(f) // esta función me retorna un slice de bytes y un error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs)) // imprimo el slice de bytes
}