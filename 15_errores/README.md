<h1>Errores</h1>

<h2>¿Por qué GO no maneja excepciones?</h2>

<p>Acoplar excepciones a una estructura de control, como la expresión try-catch-finally, da como resultado un código complicado. También tiende a incentivar a los programadores a etiquetar demasiados errores comunes, como por ejemplo no abrir un archivo, como excepcionales.</p>

<p>Go toma un enfoque diferente. Para el manejo simple de errores, los retornos multi-valor de las funciones en Go facilitan el reporte de un error sin sobrecargar el valor de retorno. <a href="https://go.dev/blog/error-handling-and-go">Un tipo de error canónico, junto con otras características de Go,</a> hace que el manejo de errores sea agradable, pero bastante diferente del de otros lenguajes.</p>

<p>Go también tiene algunas funciones integradas para señalar y recuperarse de condiciones realmente excepcionales. El mecanismo de recuperación se ejecuta solamente como parte del estado de una función que es generado después de un error, el cual es suficiente para manejar una catástrofe pero no requiere estructuras de control adicionales y, cuando se usa bien, para dar como resultado un código limpio de manejo de errores</p>


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

<h2>Printing & logging</h2>

<p>Hay varias opciones a escoger cuando necesitas imprimir o hacer logging de un error :</p>
<code>fmt.Println()</code>
<p>Imprime el error en pantalla</p>

<code>log.Println()</code>
<p>Similar a la de fmt, solo que con log podemos especificar que se guarde ese mensaje en un archivo</p>

<code>log.Fatalln()</code>
<p>Es una función donde todo muere</p>

<code>os.Exit()</code>

<code>log.Panicln()</code>
<p>Esta es un poco más suave que Fatalln() porque con esta las funciones diferidas si corren y se puede usar recover</p>

<code>panic()</code>


<p> Información tomada de: </p>
<a href="http://goconejemplos.com/panic">Enlace aquí.</a>
<a href="http://memoriascimted.com/wp-content/uploads/2021/08/Programacion-estructurada-en-Go-lang.pdf">Segundo enlace aquí.</a>