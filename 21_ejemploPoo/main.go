package main

import (
	"fmt"
	"f/models"
)

func main() {
	emp := models.Empleado{}
	emp.ConstructorHumano(1, "Yeraldine", 19)
	emp.ConstructorEmp(500000)
	
	fmt.Println(emp)
}

// cambiar a carpeta con nombre sin guion, mayusculas etcera para que funcione
// imprime -> {{1 Yeraldine 19} 500000}