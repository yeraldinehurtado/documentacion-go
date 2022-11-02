package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main(){

	//hacemos la variable de waitgroup
	wg := &sync.WaitGroup{}

	// creamos channel
	IDsChan := make(chan string)

	// cuantas goroutines vamos a aesperar
	wg.Add(2)

	//creamos las dos goroutines
	go generateIDS(wg, IDsChan) //pasamos nuestro channel
	go logIDs(wg, IDsChan)

	wg.Wait()

}

func generateIDS(wg *sync.WaitGroup, idsChan chan string){ // creamos goroutine
	// para enviar multiples valores de foma concurrente

	for i := 0; i < 100; i++ {
		id := uuid.New() // generamos un id

		//este valor que estamos recibiendo aqui, lo vamos a pasar al
		//channel
		//idsChan <- id.String() // este id lo convertimos a tipo string
		idsChan <- fmt.Sprintf("%d, %s", i+1, id.String()) // vamos a mostrar indice + el id
	}
	

	//cerramos nuestro channel para que no siga esperando valores
	close(idsChan)


	
	wg.Done()
}

func logIDs(wg *sync.WaitGroup, idsChan chan string) { // creamos una nueva goroutine
	//id := <- idsChan // asi indicamos que el valor tipo string
	// que viene dentro del channel va a ser trasladado a 
	// nuestra nueva variable id

	for id := range idsChan { //vamos a recorrer todos los valores que vengan en el channel y se lo vamos a pasar a id
		fmt.Println(id)
	}

	//fmt.Println(id) // mostramos el id

	wg.Done()
	
}