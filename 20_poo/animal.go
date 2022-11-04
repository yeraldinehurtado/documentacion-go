package main

import "fmt"

// Go no tiene clases pero si usa structs
type Animal struct {
	Id   int
	Name string
}


// Ahora creo un constructor que en go se haría por medio de
// una función
func NewAnimal(id int, name string) *Animal { // Un * previo al struct
	// o algún tipo de dato nos permite indicar donde vamos a almacenar
	// la información que este caso vamos a devolver

	return &Animal{ // un & previo al struct nos indica que vamos a
		// devolver la referencia o dirección en memoria de esa nueva
		//variable creada
		Id:   id,
		Name: name,
	}
}

// Ahora podemos agregar comportamientos para editar el nombre
// del objeto
func (e *Animal) setName(name string) {
	e.Name = name
}

func (e *Animal) GetName() string {
	return fmt.Sprintf("This is my name %s", e.Name)
}

// -----------------

//polimorfismo
func(d Animal) TellMeSomething() string {
	return "Sound"
}

// ---------

// Herencia
type Dog struct {
	Animal
}

func NewDog (id int, name string) *Dog {
	return &Dog{
		Animal{
			id,
			name,
		},
	}
}

func(d Dog) TellMeSomething() string {
	return "guau"
}

// ------------

//polimorfismo
type TellMeSomethingInterface interface {
	TellMeSomething() string
}

func TellMeSomething(g TellMeSomethingInterface) string {
	return g.TellMeSomething()
}

func main() {
	dog := NewDog(1, "Dog Animal")
	animal := NewAnimal(2, "Animal")

	fmt.Println(TellMeSomething(animal))
	//sound

	fmt.Println(TellMeSomething(dog))
	//Guau
}
