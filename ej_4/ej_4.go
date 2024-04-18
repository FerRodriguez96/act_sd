package main

import (
	"fmt"
	"sync"
)
/*
	Desarrolle un programa que tenga una variable global x iniciada en 0 (cero) y una función
	incrementar() que incremente en 5 a la variable x. La gorutina principal debe lanzar 100
	gorutinas que invoquen a la función incrementar() y luego imprimir el valor de x. Ejecute su
	programa usando la bandera -race para detectar si hay una carrera de datos. Además, el valor
	final de x debe ser 500, pero es posible que observe que a veces es 490 o 495 u otros
	valores. Usando WaitGroup y Mutexes, corrija su programa para que imprima el valor
	correcto y no tenga una carrera de datos.
*/
// uso de -race cuando se quiere detectar una carrera de datos

var x = 0;
var wg sync.WaitGroup
var cerrojo sync.Mutex

func incrementar() {
	cerrojo.Lock()
	x = x + 5
	cerrojo.Unlock()
	wg.Done()
}

func main(){
	i := 0
	for i < 100{
		wg.Add(1)
		go incrementar()
		i++
	}
	wg.Wait()
	fmt.Println(x);

}