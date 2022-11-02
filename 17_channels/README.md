<h1>Channels</h1>

<h2>Declaración e invocación</h2>

<u>
<li>Para crear un canal</li>
<code>c := make(chan int)</code>

<li>Colocando un valor en un canal</li>
<code>c <- 42</code>

<li>Tomando un valor de un canal</li>
<code><- c</code>

<li>Buffered channels</li>
<code>c := make(chan int, 4)</code>
</u>



<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>
<a href="http://memoriascimted.com/wp-content/uploads/2021/08/Programacion-estructurada-en-Go-lang.pdf">Segundo enlace aquí.</a>