<h1>Errores</h1>

<h2>Errors</h2>

<p>En Go es idiomatico comunicar errores a través de un valor de retorno separado. Esto contrasta con las excepciones usadas en lenguajes como Java y Ruby y sobrecargar el resultado como en ocasiones se hace en C. Con la manera en que se hace en Go es más facil ver cuales funciones regresan errores y manejarlos utilizando las mismas estructuras de control como lo hacemos con todas las demás tareas.</p>

<p>Por convención, los errores siempre son el último valor que se regresa y son de tipo error, una interfaz que es parte del lenguaje.</p>

<p>errors.New construye un valor de error básico con el mensaje proporcionado.</p>

```
func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}

```

<h2>Panic</h2>

<p>Una llamada a panic típicamente significa que sucedió un error inesperado. Lo usamos principalmente para terminar la ejecución en caso de errores que no debieran aparecer durante la operación normal o que no estamos listos para manejar adecuadamente.</p>

<p>Un uso común de panic es abortar la ejecución si una función devuelve un valor de error que no sabemos (o queremos) manejar. Este es un ejemplo de panic si encontramos un error inesperado al momento de crear un archivo nuevo</p>

```
package main

import "os"

func main() {
    panic("un problema")
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}

```




<p> Información tomada de: </p>
<a href="http://goconejemplos.com/panic">Enlace aquí.</a>
<a href="http://memoriascimted.com/wp-content/uploads/2021/08/Programacion-estructurada-en-Go-lang.pdf">Segundo enlace aquí.</a>