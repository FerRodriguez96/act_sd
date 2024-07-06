package main

import (
	db "db/pkg"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// TODO: Implemente el manejo de argumentos
	addr := flag.String("addr", "localhost:8000", "la dirección del servidor en formato host:port")
	flag.Parse()

	servidorReal := grpc.NewServer()
	db.RegisterBaseServer(servidorReal, db.NuevoServidor())
	// TODO: completar para usar argumentos
	listen, err := net.Listen("tcp", *addr) 

	if err == nil {
		log.Printf("Servidor escuchando en puerto %s", *addr )
	}

	if err != nil {
		log.Fatalf("fallo al escuchar: %v", err)
	}
	
	if err := servidorReal.Serve(listen); err != nil {
		log.Fatalf("fallo al servir: %v", err)
	}
}
