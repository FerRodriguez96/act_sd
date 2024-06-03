package main

/* 
	Realice un programa que cree tres canales, pase los valores "Hola", "Mundo" y "!" a los
 	canales 1, 2 y 3 respectivamente, y luego reciba datos de los canales en una declaración de
 	selección (select) e imprima lo que recibe en cada canal. Nota: es posible que las líneas de
 	impresión no estén en el orden "¡Hola, mundo!" porque el orden de impresión depende del
 	orden que van terminando las gorutinas.
*/

import (
	"fmt"
)

func main(){

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func (){
		ch1 <- "Hola"
	}()

	go func (){
		ch2 <- "Mundo"
	}()

	go func (){
		ch3 <- "!"
	}()


	for i := 0; i < 3; i++ {
		select {
			case msg1 := <-ch1:
				fmt.Println("Canal 1:", msg1)
			case msg2 := <-ch2:
				fmt.Println("Canal 2:", msg2)
			case msg3 := <-ch3:
				fmt.Println("Canal 3:", msg3)
		}
	}
}