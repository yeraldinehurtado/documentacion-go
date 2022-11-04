<h1>Estructura inicial de un programa</h1>

<p>Todo programa en Go comienza con la sentencia <em>package.</em> En el archivo principal (main) la primera línea es <em>package main.</em></p>

<p>Todo paquete principal debe tener una función principal, llamada <em>main(). </em>El contenido de la función <em>main()</em> debe ir entre llaves. Para indicar que <em>main()</em> es una función utilizamos la palabra clave (keyword) <strong>func.</strong></p>

```
package main

func main(){

}

```

<h2>Las palabras reservadas en GO</h2>

<p>En GO existen 25 palabras reservadas, que <strong>no se pueden usar como nombres de variables. </strong>Estas palabras son:</p>

<ul>
    <li>break</li>
    <li>case</li>
    <li>chan</li>
    <li>const</li>
    <li>continue</li>
    <li>default</li>
    <li>defer</li>
    <li>else</li>
    <li>fallthrough</li>
    <li>for</li>
    <li>func</li>
    <li>go</li>
    <li>goto</li>
    <li>if</li>
    <li>import</li>
    <li>interface</li>
    <li>map</li>
    <li>package</li>
    <li>range</li>
    <li>return</li>
    <li>select</li>
    <li>struct</li>
    <li>switch</li>
    <li>type</li>
    <li>var</li>
</ul>

<h2>Imprimir por pantalla.</h2>

<h3>Función Println()</h3>

<p>Si queremos imprimir una frase por pantalla podemos utilizar la función Println() contenida en el paquete fmt. Esta función debe ir dentro de la función main(). La frase que queramos mostrar por pantalla debe ir entre comillas dobles y dentro del paréntesis de la función Println(). Además si queremos trabajar con la función Println() debemos importar el paquete que la contiene, llamado fmt. Esto se consigue con la sentencia import y escribiendo entre comillas dobles el nombre del paquete.</p>

<p>Por ejemplo:</p>

```
package main

import "fmt"

func main() {
    fmt.Println("Hola mundo")
}
```

<h3>Función Print()</h3>

<p>La función Print() es muy similar a la función Println(), pero con la diferencia de que no crea un salto de línea después de ejecutarse. Se suele usar mucho menos que la función Println().</p>

<p>Por ejemplo:</p>

```
package main

import "fmt"

func main() {
    fmt.Print("Primera línea.")
    fmt.Print("Me temo que también primera línea.")
}
```

<h3>Función Printf()</h3>

<p>La función Printf() permite tener mucho más control sobre la salida. En general la función Printf() sustituye cada una de las apariciones de %v por el valor que colocamos después de la cadena de texto.</p>

<p>Por ejemplo:</p>

```
fmt.Printf("Un salto de línea.\n")

fmt.Printf("La capital de España es %v.\n", "Madrid")

fmt.Printf("Me llamo %v y tengo %v años.", "Rambo", 100)

```

<h3>Función Sprintf()</h3>

<ul>
    <li><strong>%s</strong>: Es para los String</li>
    <li><strong>%d</strong>: Es para los números ya sean enteros, flotantes, etcétera.</li>
    <li><stong>%t</stong>: para booleanos</li>
    <li><strong>%v</strong>: Cuando no se sabe que valor retorna.</li>
    <li><strong>%p</strong>: para punteros y canales.</li>
    <li><strong>%g</strong>: Para float32, complex64.</li>
</ul>

<h4>Información tomada de: </h4>

<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>