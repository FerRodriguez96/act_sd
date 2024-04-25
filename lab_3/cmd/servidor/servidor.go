package main

import (
	"context"
	"log"
	contador "contador/pkg"
	"net"
	"google.golang.org/grpc"
)

// un nuevo tipo que implementa mensajero.UnimplementedContadorServer
type servidor struct {
	contador.UnimplementedContadorServer
}

func main() {
	// TODO: Implemente el inicio del servidor
	// Use NuevoServidor() definido en servidor_nucleo.go

	

}
