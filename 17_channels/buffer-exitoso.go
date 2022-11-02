package main

import "fmt"

func main4(){
	// con buffer - buffered channel
	ca := make(chan int,1) // el segundo elemnto es el tamaño
	//del buffer y el buffer es la cantidad de elemntos o valores
	// del tipo del primer elemento (tipo int) que se pueden quedar
	//allí en el canal

	ca <- 42 // ya no necesitamos goroutine

	fmt.Println(<-ca) 
}