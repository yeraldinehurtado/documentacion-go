<h1>Punteros</h1>

<p>Son un tipo de dato que almacena la dirección de memoria de una variable. La dirección mencionada hace referencia a la ubicación del dato dentro de la memoria principal en la cual se ejecuta el programa.</p>

<p>Como Go es un lenguaje estáticamente tipado las variables puntero tienen que ser de un tipo de dato específico.
Para crear un puntero se utiliza el operador ‘*’ antes del tipo de dato que necesitamos almacenar en esa dirección de memoria</p>

<code>var x *int</code>

<h2>Cuándo usar punteros</h2>

<p>Los punteros permiten compartir un valor almacenado en alguna ubicación de la memoria. Usa pointers cuando: </p>
<ol>
<li>No quiere pasar una cantidad grande de datos</li>
<li>Quieres cambiar los datos en esa ubicación.</li>
</ol>



<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>
<a href="http://memoriascimted.com/wp-content/uploads/2021/08/Programacion-estructurada-en-Go-lang.pdf">Segundo enlace aquí.</a>