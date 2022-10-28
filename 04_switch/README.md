<h1>Switch</h1>

<p>La sentencia <em>switch</em> se puede usar con expresiones (en general será una variable) o también se puede usar sin una expresión que vaya con switch. Después de la sentencia pueden venir varios <em>case.</em> En cada case se realiza una comprobación, que debe finalizar con dos puntos. Posteriormente vienen las sentencias que se ejecutan en el caso de que la comprobación sea verdadera. Existe el caso <em>dafult</em> (que es optativo) y que se ejecuta si ninguno de los anteriores casos se cumplen. Solamente se ejecuta el primer caso que cumple la condición.</p>

<p>Sintaxis: </p>
```
switch expresión {
    case 1:
        bloque 1
    case 2:
        bloque 2
    case 3:
        bloque 3
    default:
        bloque 4
}
```

<p> Información tomada de: </p>
<a href="https://awebytes.files.wordpress.com/2020/10/librov1.pdf">Enlace aquí.</a>