<h1>Defer</h1>

<p>Una declaración "defer" (diferir) invoca a una función cuya ejecución es diferida para el momento en el que la función donde está contenida retorna, en cualquiera de los siguientes casos: o la función ejecuta una  declaración de retorno, o llega al final de su cuerpo o porque la gorutina (goroutine) correspondiente está en pánico (panicking).</p>

<p>En pocas palabras, diferir una función no es más que agarrar y decirle a una función que se ejecute al final de la función que la contiene.</p>