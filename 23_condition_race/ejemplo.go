package main

import (
	"fmt"
	"runtime"
	"sync"
)

// habilitar detección de data race en mi rpograma se hace con -race
// go run -race ejemplo.go 

func main() {
	fmt.Println("Número de CPUs", runtime.NumCPU())
	fmt.Println("Número de Goroutine", runtime.NumGoroutine())
	contador := 0

	const goroutines = 100 // el número de goroutines que voy a lanzar de manera independiente

	var wg sync.WaitGroup

	//setear nuestro contador
	wg.Add(goroutines) // Numero de goroutines

	for i := 0; i < goroutines; i++ {
		go func() {
			v := contador     // v va  a leer sobre contador
			v++               // incremento a v en una unidad
			runtime.Gosched() //cede el procesador permitiendo a otras goroutines correr
			contador = v      // escribo sobre contador

			// cada vez que una goroutine finalice su ejecución se ejecuta Done y le resta 1 al contador
			// quien monitorea el contador : wait
			wg.Done()

		}()
		fmt.Println("Número de Goroutine", runtime.NumGoroutine())
	}
	wg.Wait()                          // espera a que todas las goroutines finalicen
	fmt.Println("Contador:", contador) //imprimimos contador
}

// data race: es cuando tenemos varias goroutinas accediendo a una
// variable, comparten el acceso a esa variable y acceden simultáneamente