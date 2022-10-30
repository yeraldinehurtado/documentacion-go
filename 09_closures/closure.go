package main

import (
	"fmt"
	"strings"
)

// Los closures son funciones dentro de otra función

func repeat(num int) func(cadena string) string{
	return func(cadena string) string {
		return strings.Repeat(cadena, num)
	}

}

func main2(){
	repetir := repeat(5)
	fmt.Println(repetir("a")) // Imprime -> aaaaa
}