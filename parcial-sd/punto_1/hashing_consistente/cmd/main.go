package main

import (
	"fmt"
	pkg "parcial/pkg"
)

func main () {
	// se agregan nodos
	pkg.AgregarNodo("nodo1")
	pkg.AgregarNodo("nodo2")
	pkg.AgregarNodo("nodo3")
	// se agregan datos
	pkg.AgregarDato("dato1")
	pkg.AgregarDato("dato2")
	pkg.AgregarDato("dato3")
	// se imprime el sucesor de "dato1"
	fmt.Println("dato1 en " + pkg.GetSucesor("dato1"))
	// se imprimen los nodos y los datos
	fmt.Println(pkg.GetNodos())
	fmt.Println(pkg.GetDatos())
}