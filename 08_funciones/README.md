<h1>Funciones</h1>

<h2>Declaración e invocación</h2>

<p>Una función es un trozo de código al que le podemos dar un nombre y que es capaz de ejecutar una cierta acción. El código de una función se declara una única vez y la función puede ser llamada o invocada todas las veces que queramos. Estructurar un código a base de funciones sirve para reutilizar el código.</p>

<p>Las funciones se declaran fuera de la función <em>main()</em> utilizando la palabra reservada <em>func.</em> El llamado cuerpo de la función se escribe entre llaves y serán una o varias instrucciones que se ejecutarán cada vez que sea llamada la función. Para invocar una función debemos escribir su nombre, seguido de paréntesis.</p>

<p>Sintaxis de una función:</p>

```
func nombreFuncion(){

}
```

<h2>Funciones con parámetros</h2>

<p>Las funciones anteriores siempre realizan exactamente la misma acción. Podemos incluir parámetros en la función de tal forma que la ejecución de la función no sea siempre exactamente
igual, sino que dependa del parámetro o parámetros que les pasemos. Los parámetros se escriben dentro de los paréntesis de la función y tienen que tener un nombre y un tipo. En realidad los parámetros son nombres de variables que se pueden usar en el cuerpo de la función.</p>

```
func nombreFuncion(nVariable tipoDato){

}

```

<p>Por ejemplo:</p>

```
func saludar2(nombre string) {
    fmt.Println("Hola, soy", nombre)
}
```

<h2>Funciones con retorno</h2>

<p>Hasta ahora las funciones que hemos creado ejecutan una
acción, pero no devuelven un valor que podamos, por ejemplo,
almacenar en una variable. Si queremos que esto ocurra
debemos decirle a la función qué tipo de valor nos debe
devolver. Para ello utilizamos la palabra clave return antes
del valor que queremos devolver. También es obligatorio
indicar al compilador el tipo de dato que vamos a devolver.
Este tipo se escribe antes de la apertura de llaves. El tipo de la función también se hace eco del retorno de la función.</p>

<p>Por ejemplo:</p>

```
func area(base int, altura int) int {
    resultado := base * altura
    return resultado
}

```

<h2>Funciones con retorno múltiple</h2>

<p>En realidad en vez de un valor, una función también puede
devolver más de un valor. Dichos valores deben ir escritos
después de la palabra return. Como devolvemos más de un valor debemos escribir dichos tipos, entre paréntesis, en la
declaración de la función.</p>

<p>Por ejemplo:</p>

```
func sumaResta(x int, y int) (int, int) {
    return x + y, x-y
}

```

<h2>Funciones con retorno nombrado</h2>

<p>A veces nos interesa devolver el valor de una variable que
creamos dentro del cuerpo de la función. Para este caso especial podemos utilizar los retornos nombrados. En la definición de función escribimos el nombre y el tipo de la variable que vamos a retornar. De esta forma basta con añadir la sentencia return sin añadir nada más.</p>

<p>Por ejemplo:</p>

```
func h2(a int, b int) (r int) {
    r = a * a + b * b
    return // No añadimos nada al return
}

```


<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>
<a href="http://memoriascimted.com/wp-content/uploads/2021/08/Programacion-estructurada-en-Go-lang.pdf">Segundo enlace aquí.</a>