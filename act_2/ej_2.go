package main

/*
Lea del apunte anexo a la clase sobre WaitGroup y úselo en el siguiente problema. Cree una
función tarea que reciba un id y tiempo que dura la tarea, en la función se debe imprimir el
id de la tarea, cuánto tiempo dura, esperar el tiempo especificado, e imprimir que va a
finalizar. La gorutina principal debe lanzar 5 tareas de forma concurrente, la gorutina debe
esperar por la finalización de las tareas lanzadas e imprimir que han finalizado.

 se debe invocar a wg.add antes de llamar a la rutina, y poner el wg.wait al final
*/

import "fmt"

func main() {
	fmt.Println("Hola mundo")
}