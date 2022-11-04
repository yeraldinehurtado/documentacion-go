package main

import (
	"fmt"
	"time"
)

func trabajo(terminado chan bool) {
	fmt.Print("Trabajando...")
	time.Sleep(time.Second)
	fmt.Println("terminado")

	terminado <- true
}

func main() {
	terminado :=  make(chan bool, 1)
	go trabajo(terminado)

	<-terminado
}