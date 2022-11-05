package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizado")
}

// ENviar
func enviar(p, i, s chan<- int) { // send only
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j // envio al canal par lo que hay en j
		} else {
			i <- j // sino enviamos el impar a i
		}
	}

	s <- 0 // enviamos 0 a salir
}

// REcibir
func recibir(p, i, s <-chan int) { // receive only
	for { // for infinito
		select {
		case v := <-p: // si hay un valor en el canal par
			fmt.Println("Desde el canal par:", v)
		case v := <-i: // si hay un valor en el canal impar
			fmt.Println("Desde el canal impar:", v)
		case v := <-s: // si hay un valor 0 en el canal salir
			fmt.Println("Desde el canal salir:", v)
			return
		}
	}
}

/*
Imprime->

Desde el canal par: 0
Desde el canal impar: 1
Desde el canal par: 2
Desde el canal impar: 3
Desde el canal par: 4
Desde el canal impar: 5
Desde el canal par: 6
Desde el canal impar: 7
Desde el canal par: 8
Desde el canal impar: 9
Desde el canal par: 10
Desde el canal impar: 11
Desde el canal par: 12
Desde el canal impar: 13
Desde el canal par: 14
Desde el canal impar: 15
Desde el canal par: 16
Desde el canal impar: 17
Desde el canal par: 18
Desde el canal impar: 19
Desde el canal par: 20
Desde el canal impar: 21
Desde el canal par: 22
Desde el canal impar: 23
Desde el canal par: 24
Desde el canal impar: 25
Desde el canal par: 26
Desde el canal impar: 27
Desde el canal par: 28
Desde el canal impar: 29
Desde el canal par: 30
Desde el canal impar: 31
Desde el canal par: 32
Desde el canal impar: 33
Desde el canal par: 34
Desde el canal impar: 35
Desde el canal par: 36
Desde el canal impar: 37
Desde el canal par: 38
Desde el canal impar: 39
Desde el canal par: 40
Desde el canal impar: 41
Desde el canal par: 42
Desde el canal impar: 43
Desde el canal par: 44
Desde el canal impar: 45
Desde el canal par: 46
Desde el canal impar: 47
Desde el canal par: 48
Desde el canal impar: 49
Desde el canal par: 50
Desde el canal impar: 51
Desde el canal par: 52
Desde el canal impar: 53
Desde el canal par: 54
Desde el canal impar: 55
Desde el canal par: 56
Desde el canal impar: 57
Desde el canal par: 58
Desde el canal impar: 59
Desde el canal par: 60
Desde el canal impar: 61
Desde el canal par: 62
Desde el canal impar: 63
Desde el canal par: 64
Desde el canal impar: 65
Desde el canal par: 66
Desde el canal impar: 67
Desde el canal par: 68
Desde el canal impar: 69
Desde el canal par: 70
Desde el canal impar: 71
Desde el canal par: 72
Desde el canal impar: 73
Desde el canal par: 74
Desde el canal impar: 75
Desde el canal par: 76
Desde el canal impar: 77
Desde el canal par: 78
Desde el canal impar: 79
Desde el canal par: 80
Desde el canal impar: 81
Desde el canal par: 82
Desde el canal impar: 83
Desde el canal par: 84
Desde el canal impar: 85
Desde el canal par: 86
Desde el canal impar: 87
Desde el canal par: 88
Desde el canal impar: 89
Desde el canal par: 90
Desde el canal impar: 91
Desde el canal par: 92
Desde el canal impar: 93
Desde el canal par: 94
Desde el canal impar: 95
Desde el canal par: 96
Desde el canal impar: 97
Desde el canal par: 98
Desde el canal impar: 99
Desde el canal salir: 0
Finalizado
*/
