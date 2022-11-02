package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main(){

	//hacemos el waitgroup
	wg := &sync.WaitGroup{}

	// creamos channel
	IDsChan := make(chan string)

	//creamos las dos goroutines
	go generateIDS(wg, IDsChan) //pasamos nuestro channel
	go logIDs(wg, IDsChan)

	wg.Wait()

}

func generateIDS(wg *sync.WaitGroup, idsChan chan string){ // creamos goroutine
	id := uuid.New() // generamos un id

	//este valor que estamos recibiendo aqui, lo vamos a pasar al
	//channel
	idsChan <- id.String() // este id lo convertimos a tipo string

	wg.Done()
}

func logIDs(wg *sync.WaitGroup, idsChan chan string) { // creamos una nueva goroutine
	id := <- idsChan // asi indicamos que el valor tipo string
	// que viene dentro del channel va a ser trasladado a 
	// nuestra nueva variable id

	fmt.Println(id) // mostramos el id
	
}