package main

import "fmt"

func main() {
	mensajes := make(chan string, 2)

	mensajes <- "Buffered"
	mensajes <- "Channel"

	fmt.Println(<-mensajes)
	fmt.Println(<-mensajes)
}

/*
imprime ->
Buffered
Channel
*/