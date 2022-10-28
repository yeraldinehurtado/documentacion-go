package main 

import "fmt"

func main5(){
	//Los valores que no se facilitan GO los inicializa en 0

	var a = [...]int{2: 2222, 4: 8888, 1: 1234}
	fmt.Println(a) // -> [0 1234 2222 0 8888]
}