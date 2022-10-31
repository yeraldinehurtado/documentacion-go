package main

import "fmt"

type rectangulo struct {
	ancho, alto int
}

// func (r receptor) identificador(parametros) (retorno(s)) { codigo }

func (r *rectangulo) area() int { // apunta a rectangulo y retorna un int
	return r.ancho * r.alto 
}

func (r *rectangulo) perimetro() int {
	return 2 * r.ancho + 2*r.alto
}

func main2() {
	r := rectangulo{
		ancho: 10,
		alto: 5,
	}

	fmt.Println("area: ", r.area())
	fmt.Println("Perimetro: ", r.perimetro())

}