package main

import "fmt"

func main2(){
	// sin buffer - unbuffered channel
	ca := make(chan int)

	ca <- 42 // enviando entero 42
	fmt.Println(<-ca)
}

// esta es una gorutine que envia, y se va a bloquear porque no
// hay otra gorutina que reciba


// imprime -> fatal error: all goroutines are asleep - deadlock!