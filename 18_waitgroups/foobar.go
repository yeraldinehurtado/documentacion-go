package main

import (
	"fmt"
	"runtime"
	"sync"
)

// declaramos una variable a nivel de paquete de tipo sync.WaitGroup
var wg sync.WaitGroup // con esta variable podremos tener acceso a los metodos
// done, wait y add

func main() {
	fmt.Println("OS\t\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	// antes de llamar a las goroutines , utilizar el
	// metodo add para setear un contador que va a estar
	// monitoreando cuantas goroutines va a haber en mi programa
	// y cuntas van terminando, para que al final diga que
	// "hey, ya no hay mas goroutines, puedes terminar."

	// seteamos nuestro contador con el número de goroutines
	//que va a haber
	wg.Add(1) // el número de goroutines

	go foo()
	bar()

	// wait monitorea ese contador y cuando llegue a 0
	// le dice a main que ya no hay goroutines por ejecutar,
	// puede finalizar
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() //cuando termina le resta al contador que se setea
	// con add(1) entonces tenemos una goroutine menos
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

/*
imprime ->

OS			 linux
ARCH		 amd64
CPUs		 8
Gorutines	 1
bar: 0
bar: 1
bar: 2
bar: 3
bar: 4
bar: 5
bar: 6
bar: 7
bar: 8
bar: 9
foo: 0
foo: 1
foo: 2
foo: 3
foo: 4
foo: 5
foo: 6
foo: 7
foo: 8
foo: 9

*/
