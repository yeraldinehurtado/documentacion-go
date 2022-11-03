package main

import "fmt"

type Person struct{
	name string
}

func (p Person) Name() string {
	return p.name
}

func (p *Person) setName(name string) {
	p.name = name
}

func main(){
	me := Person{"Yeraldine"}
	fmt.Println("Mi nombre es: ", me.Name())

	me.setName("Juana")
	fmt.Println("Mi nuevo nombre es: ", me.Name())
}