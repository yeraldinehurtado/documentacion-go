<h1>MAPS</h1>

<h2>Concepto de map</h2>

<p>Los mapas en Go son una generalización de los Array. Son conjuntos de pares clave/valor. Todas las claves en un mapa deben pertenecer a un tipo determinado. Los valores también tienen que ser también todos del mismo tipo. Para crear un map debemos informar al compilador del tipo de dato de la clave y del tipo de dato del valor. Los pares de valores se separan por comas y entre la clave y el valor debemos colocar dos puntos.
Para mejorar la lectura del código cada par se escribe en una línea, no olvidando escribir la coma final.</p>

<p>La sintaxis general de un mapa es: </p>

<code>map[tipoClave]tipoValor</code>

<p>Si el mapa tiene pocos datos se puede escribir en una única línea</p>

<code>n := map[string]bool{"a" : true, "adios" : false}</code>

<p>Los mapas pueden crecer de la siguiente forma: </p>

```
m := map[string]int{
    "Euler": 1707,
    "Riemann": 1826,
}
m["Hilbert"] = 1862
fmt.Println(m)
```
<p>Imprime -> map[Euler:1707 Hilbert:1862 Riemann:1826]</p>

<p>El orden de las parejas no es importante en los mapas.</p>

<h2>Para borrar pares...</h2>

<p>Podemos borrar pares con la función <em>delete().</em> Debemos facilitar el nombre del mapa y también la clave que queremos borrar.</p>
 
```
m := map[string]int{
    "Euler": 1707,
    "Riemann": 1826,
}
delete(m, "Euler")
fmt.Println(m) // -> map[Riemann:1826]
```


<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>
<a href="http://memoriascimted.com/wp-content/uploads/2021/08/Programacion-estructurada-en-Go-lang.pdf">Segundo enlace aquí.</a>