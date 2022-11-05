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


/*
Lo que imprime es muy inconsistente, imprime cosas diferentes
una de esas es esta ->

Número de CPUs 8
Número de Goroutine 1
Número de Goroutine 2
Número de Goroutine 3
Número de Goroutine 4
Número de Goroutine 5
Número de Goroutine 6
Número de Goroutine 7
Número de Goroutine 8
Número de Goroutine 9
Número de Goroutine 10
Número de Goroutine 11
Número de Goroutine 12
Número de Goroutine 13
Número de Goroutine 14
Número de Goroutine 15
Número de Goroutine 16
Número de Goroutine 17
Número de Goroutine 18
Número de Goroutine 17
Número de Goroutine 18
Número de Goroutine 16
Número de Goroutine 12
Número de Goroutine 11
Número de Goroutine 9
Número de Goroutine 9
Número de Goroutine 8
Número de Goroutine 5
Número de Goroutine 4
Número de Goroutine 5
Número de Goroutine 6
Número de Goroutine 7
Número de Goroutine 8
Número de Goroutine 9
Número de Goroutine 10
Número de Goroutine 11
Número de Goroutine 11
Número de Goroutine 12
Número de Goroutine 13
Número de Goroutine 14
Número de Goroutine 12
Número de Goroutine 10
Número de Goroutine 11
Número de Goroutine 12
Número de Goroutine 13
Número de Goroutine 14
Número de Goroutine 15
Número de Goroutine 16
Número de Goroutine 17
Número de Goroutine 18
Número de Goroutine 19
Número de Goroutine 20
Número de Goroutine 18
Número de Goroutine 16
Número de Goroutine 15
Número de Goroutine 16
Número de Goroutine 10
Número de Goroutine 11
Número de Goroutine 12
Número de Goroutine 12
Número de Goroutine 11
Número de Goroutine 4
Número de Goroutine 2
Número de Goroutine 3
Número de Goroutine 4
Número de Goroutine 5
Número de Goroutine 3
Número de Goroutine 4
Número de Goroutine 5
Número de Goroutine 6
Número de Goroutine 7
Número de Goroutine 8
Número de Goroutine 9
Número de Goroutine 10
Número de Goroutine 11
Número de Goroutine 12
Número de Goroutine 13
Número de Goroutine 14
Número de Goroutine 15
Número de Goroutine 16
Número de Goroutine 17
Número de Goroutine 18
Número de Goroutine 19
Número de Goroutine 20
Número de Goroutine 21
Número de Goroutine 22
Número de Goroutine 23
Número de Goroutine 24
Número de Goroutine 25
Número de Goroutine 26
Número de Goroutine 27
Número de Goroutine 28
Número de Goroutine 29
Número de Goroutine 30
Número de Goroutine 27
Número de Goroutine 28
Número de Goroutine 27
Número de Goroutine 25
Número de Goroutine 26
Número de Goroutine 27
Número de Goroutine 23
Número de Goroutine 17
Contador: 15

*/