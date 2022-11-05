package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go mi_nombre_lento("Yeraldine") // goroutine para que sea concurrente
	fmt.Println("Hola")
	var wait string
	fmt.Scanln(&wait) // esto es para que no termine el programa hasta que yo coloque un enter

}

func mi_nombre_lento(name string) {
	letras := strings.Split(name, "") //split permite partir la cadena en otras cadenas

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) // que duerma un segundo cuando trata de imprimir la siguiente letra
		fmt.Println(letra)
	}

}

// Nota : Todas las goroutines deben tener el tiempo suficiente para poder completarse
// Si se alcanza el final del script o se cierra la ejecución manualmente las rutinas
// en proceso también se cancelarán, ahí entran los canales


// concurrencia significa dividir un problema en diferentes
//ejecuciones que se ejecutan de forma concurrente (al mismo tiempo)
