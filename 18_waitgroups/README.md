<h1>WaitGroup</h1>

<p><strong>Un WaitGroup espera que una colección de gorutinas terminen su trabajo.</strong> La gorutina de main llama Add para configurar el número de gorutinas por las que tiene que esperar. Luego cada una de las gorutinas corre y llama a Done cuando terminan. Al mismo tiempo, Wait puede ser usado para bloquear hasta que todas las gorutinas hayan finalizado. Escribir código concurrente es súper fácil: todo lo que tenemos que hacer es <strong> poner “go” en frente de una llamada a una función o método. </strong></p>

<code>runtime.NumCPU()</code>

<code>runtime.NumGoroutine()</code>

<p>TYPE WaitGroup</p>

<code>sync.WaitGroup</code>

<code>func (wg *WaitGroup) Add(delta int) </code>

<code>func (wg *WaitGroup) Done() </code>

<code>func (wg *WaitGroup) Wait()</code>


<h4>Información tomada de: </h4>
<p>Guía del curso - Eduar Tua</p>