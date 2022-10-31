<h1>Interfaces y polimorfismo</h1>

<p>En GO, los valores pueden ser de más de un tipo. Una interfaz permite a un valor ser de más de un tipo. Creamos una interfaz usando la sintaxis: </p>

```
type humano interface

```

<p>Luego definimos cuáles métodos debe tener un tipo para implementar esa interfaz. Si un tipo tiene los métodos requeridos, el cuál podrían ser ninguno, (la interfaz vacía denotada por interface{}), entonces ese TIPO implícitamente implementa la interfaz y es también de ese tipo de interfaz.</p>