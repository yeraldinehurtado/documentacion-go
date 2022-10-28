<h1>Slices</h1>

<p>A efectos prácticos un slice es un array cuyo tamaño puede
variar. Por lo tanto para construir un slice únicamente tenemos
que facilitarle al compilador el tipo de los elementos que va a
contener. La inicialización es muy similar a los arrays, pero
omitiendo el número de elementos.</p>

<code>s := []int{10, 20, 30}</code>

<h2>Características de los slices</h2>

<p>A partir de su funcionamiento como un puntero que almacena la dirección de los datos (de un solo dato accede al resto de elementos en adelante) de un arreglo y que a su vez tiene la capacidad de modificar los datos de este último, se permite tener cierto nivel de flexibilidad en el volumen de datos que se manejan, además de procesos como la inserción y copia, tornando su uso mucho más eficiente.</p>

<ul>
    <li><strong>LEN(): </strong>: Esta función permite conocer el tamaño de un slice (y también de un array), siendo este tamaño igual al número de elementos almacenados en la estructura, que para un slice sería el número de elementos almacenados en el array.</li>
    <li><strong>CAP(): </strong>: Esta función muestra la capacidad (cantidad de datos que puede ser almacenada) tanto para un slice como un para un array de guardar nuevos elementos.</li>
</ul>



<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>
<a href="http://memoriascimted.com/wp-content/uploads/2021/08/Programacion-estructurada-en-Go-lang.pdf">Segundo enlace aquí.</a>