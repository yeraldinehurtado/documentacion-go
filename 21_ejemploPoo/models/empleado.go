package models

type Empleado struct {
	Humano
	salario int64
}

func (e *Empleado) ConstructorEmp(salario int64) {
	e.salario = salario
}

func (e *Empleado) GetSalario() int64 {
	return e.salario
}

func (e *Empleado) SetSalario(salario int64) {
	e.salario = salario
}

// Se hace constructor cuando la estructura es pública y sus
// atributos son privados
