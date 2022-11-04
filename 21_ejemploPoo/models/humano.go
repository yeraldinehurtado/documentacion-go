package models

type Humano struct {
	id int
	nombre string
	edad int
}

func (h *Humano) ConstructorHumano(id int, nombre string, edad int) {
	h.id = id
	h.nombre = nombre
	h.edad = edad
}

func (h *Humano) GetName() string {
	return h.nombre
}

func (h *Humano) SetName(nombre string) {
	h.nombre = nombre
}

// Se hace constructor cuando la estructura es pública y sus 
// atributos son privados