<h1>Bucle for</h1>

<h2>Bucle FOR simple</h2>

<p>Un bucle es un conjunto de <strong>sentencias que pueden ejecutarse más de una vez. </strong> En GO existe un único bucle, llamado <em>for</em>, pero que puede ser usado de distintas formas. La más sencilla corresponde a la siguiente sintaxis: </p>

'''for condición {
    bloque
}'''

<p>Donde <em>condición</em> es un valor booleano, o cualquier expresión que devuelva un valor booleano. EL bloque se escribe entre llaves y se suele indentar para mejorar la lectura del código.</p>

<p>En un bucle <em>for</em> se comprueba la condición. Si esta resulta ser <em>TRUE</em> se ejecuta el código situado entre las llaves. Cuando se termina de ejecutar el bloque el flujo del programa vuelve al inicio del bucle for y todo se repite hasta que la condición se hace <em>FALSE.</em> En ese momento no se ejecuta el bloque y el flujo del programa sigue con la siguiente instrucción programada.</p>

<h3>Bucle FOR compuesto</h3>
<p>En un blucle <em>for</em> es muy habitual crear una variable que servirá de contador y que se define justamente antes de entrar en el bucle. También es muy habitual actualizar dicha variable al terminar de ejecutar todas las sentencias del bloque. Por ello los diseñadores de Go crearon otra forma de utilizar el bucle <em>for</em>, que tiene la siguiente sintaxis.</p>

'''for inicio ; condición ; actualizacion {
    bloque
}'''

<p>Esta sentencia <em>for</em> tiene tres partes:</p>
<ul>
    <li>En <em>INICIO</em> se suele declarar e inicializar una variable que solamente se puede usar dentro del bucle.</li>
    <li><em>Condición</em> es una expresión booleana. Mientras dicha condición sea true el bloque se estará ejecutando.</li>
    <li>En <em>actualización</em> se suele variar el valor de la variable que hemos definido en <em>inicio</em></li>
</ul>

<p>Esta forma de ejecutar el bucle for puede ser un poco más cómoda, pero su funcionamiento es exactamente igual al for simple.</p>

<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>