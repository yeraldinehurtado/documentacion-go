package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() { // funcion (goroutine) anónima
		time.Sleep(time.Second * 1)
		c1 <- "uno"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "dos"
	}()

	for i := 0; i < 2; i++ {
		select { // select es como switch pero para canales
		case mensaje1 := <-c1:
			fmt.Println("Recibido", mensaje1)
		case mensaje2 := <-c2:
			fmt.Println("Recibido", mensaje2)
		}
	}
}

/*
Imprime ->

Recibido uno
Recibido dos
*/
