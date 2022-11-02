package main

import "fmt"

func main(){
	// con buffer - buffered channel
	ca := make(chan int,1) // para solucionar solo se cambia el segundo valor
	//por el 2

	ca <- 42
	ca <- 43 // cuando esta línea se ejecuta , se bloquea el codigo

	fmt.Println(<-ca) 
}

// imprime -> fatal error: all goroutines are asleep - deadlock!