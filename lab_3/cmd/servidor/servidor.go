package main

import (
	"log"
	contador "contador/pkg"
)


func main() {
	// TODO: Implemente el inicio del servidor
	// Use NuevoServidor() definido en servidor_nucleo.go
	servidor := contador.NuevoServidor()
	if err := servidor.Iniciar(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
