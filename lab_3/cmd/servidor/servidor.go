package main

import (
	"log"
	contador "contador/pkg"
	"net"
	grpc "google.golang.org/grpc"
)


func main() {
	// TODO: Implemente el inicio del servidor
	// Use NuevoServidor() definido en servidor_nucleo.go

	servidorReal := grpc.NewServer()
	contador.RegisterContadorServer(servidorReal, contador.NuevoServidor())
	listen, err := net.Listen("tcp", "localhost:8000")

	if err == nil {
		log.Println("Servidor escuchando en puerto 8000")
	}


	if err != nil {
		log.Fatalf("fallo al escuchar: %v", err)
	}
	if err := servidorReal.Serve(listen); err != nil {
		log.Fatalf("fallo al servir: %v", err)
	}
}
