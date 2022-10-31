package main

import "fmt"
import "math"

type geometrica interface {
	area() float64
	perimetro() float64
}

type cuadro struct {
	ancho, alto float64
}

type circulo struct {
	radio float64
}

func (s cuadro) area() float64 {
	return s.ancho * s.alto
}

func (s cuadro) perimetro() float64 {
	return 2*s.ancho + 2*s.alto
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.radio
}

func medida(g geometrica) {
	fmt.Println(g) // imprime -> {3 4} {5}
	fmt.Println(g.area()) // imprime -> 12
	fmt.Println(g.perimetro()) // imprime -> 14
}

func main() {
	s := cuadro{ancho: 3, alto: 4}
	c := circulo{radio: 5}

	medida(s) // imprime -> 78.53981633974483
	medida(c) // imprime -> 31.41592653589793
}