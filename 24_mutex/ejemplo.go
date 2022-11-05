package main

import (
	"fmt"
	"runtime"
	"sync"
)

// esto es continuación del ejemplo anterior para evitar los data races


// data race: es cuando tenemos varias goroutinas accediendo a una
// variable, comparten el acceso a esa variable y acceden simultáneamente

// Mutex lo que hace es que bloquea el acceso tanto de
// lectura y escritura a esa variable cuando una goroutina
// esta hacinedo operaciones sobre ella

func main() {
	fmt.Println("Número de CPUs", runtime.NumCPU())
	fmt.Println("Número de Goroutine", runtime.NumGoroutine())
	contador := 0

	const goroutines = 100 // el número de goroutines que voy a lanzar de manera independiente

	var wg sync.WaitGroup

	//setear nuestro contador
	wg.Add(goroutines) // Numero de goroutines

	//mutex
	var mut sync.Mutex // con esta variable podemos usar los metodos de bloque y desbloqueo 

	for i := 0; i < goroutines; i++ {
		go func() {
			//agregamos metodo de bloqueo mutex
			mut.Lock()
			// entonces esto debería bloquear el acceso a esa variable que es compartida cada vez que una goroutina está
			// haciendo operaciones sobre ella
			


			v := contador     // v va  a leer sobre contador
			v++               // incremento a v en una unidad
			runtime.Gosched() //cede el procesador permitiendo a otras goroutines correr
			contador = v      // escribo sobre contador

			//desbloqueamos (mutex)
			mut.Unlock()

			// cada vez que una goroutine finalice su ejecución se ejecuta Done y le resta 1 al contador
			// quien monitorea el contador : wait
			wg.Done()

		}()
		fmt.Println("Número de Goroutine", runtime.NumGoroutine())
	}
	wg.Wait()                          // espera a que todas las goroutines finalicen
	fmt.Println("Contador:", contador) //imprimimos contador
}
