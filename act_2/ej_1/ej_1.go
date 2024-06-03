package main

/*
	Cree un canal sin búfer, pase los valores del 1 al 10 por al canal y luego reciba del canal e
	imprima los valores. Sugerencia: asegúrese de que está utilizando una gorutina.
*/

import "fmt"

func main() {

	// se crea un canal sin buffer de tipo int
	ch := make (chan int)

	//se llama a una gorutina
	go func () {
		//se establece un bucle
		for i:= 1; i <= 10; i++{
			//se insertan los valores que va tomando i en el canal
			ch <- i
		}
		//se cierra el canal
		close(ch)
	}()
	
	//se establece un for para consumir los datos en el rango del canal
	for datos := range ch{
		//se imprimen los datos
		fmt.Println("El valor consumido del canal es:" ,datos)
	}

}