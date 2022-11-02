package main

import "fmt"

func main3(){
	// sin buffer - unbuffered channel
	ca := make(chan int)

	go func(){ // creamos goroutine que envia el valor

		ca <- 42

	}()

	fmt.Println(<-ca) // <-- la goroutine principal va a recibir del canal
	// el valor, con esta expresion.
}