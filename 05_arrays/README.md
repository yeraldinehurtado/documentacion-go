<h1>Arrays</h1>

<p>Un array es un colección ordenada de datos del mismo tipo. Debemos informar al compilador del tipo de dato y de la cantidad de datos que contendrá. El número de datos lo puede inferir el compilador si escribimos tres puntos suspensivos en lugar del número de datos. Los datos se escriben entre llaves y separados por comas.</p>

<code>u := [5]int{10, 20, 30, 40, 50} // Tipo int y 5 datos</code>
<code>v := [...]float64{10.1, 20.1, 30.1} // Inferencia</code>

<h2>¿Cómo acceder a los datos de un array?</h2>

<p>Podemos acceder de manera individual a cada uno de los elementos del vector utilizando su posición entre corchetes. Debemos tener en cuenta que el primer elemento de cualquier vector tiene índice 0.</p>

<code>u := [5]int{10, 20, 30, 40, 50}</code>
<code>fmt.Println(u[0]) // -> 10</code>
<code>fmt.Println(u[1]) // -> 20</code>
<code>fmt.Println(u[4]) // -> 50</code>

<h2>Array sin inicializar</h2>

<p>Se puede construir un vector sin asignarle ningún valor. En este caso Go asigna por defecto el valor cero a cada uno de los items, como se puede ver en el siguiente código.</p>

<code>var vector [5]int</code>
<code>fmt.Println(vector) // -> [0 0 0 0 0]</code>

<h2>Inicialización parcial</h2>

<p>Si en un vector le facilitamos un número menor de elementos, entonces el compilador inicializa el resto al valor nulo. Si lo inicializamos con más elementos de los que puede contener entonces tendremos un error.</p>

<h2>Inicialización de modo disperso</h2>

<p>Existe otra forma de inicializar vectores, dando el índice separado por dos puntos del valor que contiene. Los valores no facilitados los inicializa Go a cero. En el siguiente ejemplo se inicializa un vector con tres componentes y de modo disperso.</p>

<code>var a = [9]int{2:2222, 8: 8888, 1: 1234}</code>
<code>fmt.Println(a)</code>

<p>Imprime -> [0 1234 2222 0 0 0 0 0 8888]</p>

<h2>Inicialización de arrays en columna</h2>

<p>Algunas veces, sobre todo si los vectores son de tipo cadena, es conveniente, para mejorar la visualización del código, escribir cada elemento en una línea. Si lo hacemos de este modo el último elemento también debe de llevar una coma.</p>

<code>var deporte = [3]string{
"canicas",
"peonza",
"chapas"
}</code>
<code>fmt.Println(deporte)</code>

<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>