package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	start := time.Now()
	wg := &sync.WaitGroup{}  // variable de tipo puntero


	// forma tradicional 
	for i := 0 ; i< 10; i++ {

		wg.Add(1) // por cada iteracion vamos a agregar 1 goroutine

		go showGoroutine(i, wg)
	}
	
	wg.Wait()
	duration := time.Since(start).Milliseconds()

	fmt.Printf("%dms\n", duration)

}

func showGoroutine(id int, wg *sync.WaitGroup){
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%d with %dms\n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}


// con la forma tradicional se tarda mas en ejecutar


/* 
for i := 0 ; i< 10; i++ {
		showGoroutine(i)
	}
con for normal imprime :

Goroutine #0 with 81ms
Goroutine #1 with 387ms
Goroutine #2 with 347ms
Goroutine #3 with 59ms
Goroutine #4 with 81ms
Goroutine #5 with 318ms
Goroutine #6 with 425ms
Goroutine #7 with 40ms
Goroutine #8 with 456ms
Goroutine #9 with 300ms

*/