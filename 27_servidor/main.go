package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Router
	http.HandleFunc("/", Hola)
	http.HandleFunc("/page", PaginaNoFunciona)
	http.HandleFunc("/saludar", Saludar)

	// Crear un servidor
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(http.ListenAndServe("localhost:3000", nil)) // ruta donde se va a ejecutar nuestro servidor

}

// Handlers
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("EL metodo es: " + r.Method) // muestra con que metodo estamos haciendo esta peticion (get, post, put...)
	fmt.Fprintln(rw, "Hola mundo")
}

// Handlers
func PaginaNoFunciona(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(rw, "Hola, %s tu edad es %s", name, age)

}

// http://localhost:3000/saludar?name=Yeraldine&age=19

/*
imprime en consola ->

El servidor está corriendo en el puerto 3000
Run server: http://localhost:3000/
/saludar

map[]
/saludar?name=Yeraldine&age=19
name=Yeraldine&age=19
map[age:[19] name:[Yeraldin
*/