package main

/*
	Lea del apunte anexo a la clase sobre WaitGroup y úselo en el siguiente problema. Cree una
	función tarea que reciba un id y tiempo que dura la tarea, en la función se debe imprimir el
	id de la tarea, cuánto tiempo dura, esperar el tiempo especificado, e imprimir que va a
	finalizar. La gorutina principal debe lanzar 5 tareas de forma concurrente, la gorutina debe
	esperar por la finalización de las tareas lanzadas e imprimir que han finalizado.
*/

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func tarea (id int, tiempo int){
	//Imprime el Id de la tarea
	fmt.Println("El id de la tarea es:", id)
	//Imprime el tiempo de espera
	fmt.Println("Y dura", tiempo, "segundos")
	//duerme el tiempo especificado
	time.Sleep(time.Duration(tiempo) * time.Second)
	//Ha finalizado la tarea!
	fmt.Println("Tarea finalizada!")
	
	wg.Done()
}

func main() {
	i := 0
	//se define un for para lanzar las rutinas 5 veces
	for i < 5{
		//se añade al grupo de espera
		wg.Add(1)
		//se lanza la tarea
		go tarea(i, 5)
		i++
	}
	//se espera a que todas las tareas finalizen
	wg.Wait()
	fmt.Println("Han finalizado todas las tareas");

}