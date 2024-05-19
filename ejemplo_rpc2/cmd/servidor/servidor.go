package main

import (
	db "db/pkg"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	// TODO: Implemente el manejo de argumentos

	servidorReal := grpc.NewServer()
	servidor := db.NuevoServidor()
	db.RegisterBaseServer(servidorReal, &servidor)
	// TODO: completar para usar argumentos
	listen, err := net.Listen("tcp", "") 
	if err != nil {
		log.Fatalf("fallo al escuchar: %v", err)
	}
	if err := servidorReal.Serve(listen); err != nil {
		log.Fatalf("fallo al servir: %v", err)
	}
}
