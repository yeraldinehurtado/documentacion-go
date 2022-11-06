package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// esto es continuación del ejemplo anterior para evitar los data races

// data race: es cuando tenemos varias goroutinas accediendo a una
// variable, comparten el acceso a esa variable y acceden simultáneamente

// go run -race ejemplo.go 

func main() {
	fmt.Println("Número de CPUs", runtime.NumCPU())
	fmt.Println("Número de Goroutine", runtime.NumGoroutine())
	var contador int64 // convertimos el contador en una variable tipo int64

	const goroutines = 100 // el número de goroutines que voy a lanzar de manera independiente

	var wg sync.WaitGroup

	//setear nuestro contador
	wg.Add(goroutines) // Numero de goroutines

	
	for i := 0; i < goroutines; i++ {
		go func() {
			
			atomic.AddInt64(&contador, 1) // del paquete atomic utilizamos addint64, y addint64 recibe como primer parametro
			// la dirección en memoria de la variable contador, y el cambio que le vamos a hacer a contador, o sea, 1 unidad

			runtime.Gosched() //cede el procesador permitiendo a otras goroutines correr

			fmt.Println("Contador: ", atomic.LoadInt64(&contador))


			// cada vez que una goroutine finalice su ejecución se ejecuta Done y le resta 1 al contador
			// quien monitorea el contador : wait
			wg.Done()

		}()
		fmt.Println("Número de Goroutine", runtime.NumGoroutine())
	}
	wg.Wait()                          // espera a que todas las goroutines finalicen
	fmt.Println("Contador:", contador) //imprimimos contador
}
