package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	db "db/pkg"
)

func main() {
	// TODO: Implemente el manejo de argumentos
	addr := flag.String("addr", "localhost:8000", "la direccion del servidor en formato host:puerto")
	esLider := flag.Bool("lider", false, "indica si el servidor es lider")
	seguidores := flag.String("seguidores", "", "lista de direcciones de seguidores separadas por coma")
	flag.Parse()

	ListaSeguidores := []string{}
	if *seguidores != "" {
		ListaSeguidores = strings.Split(*seguidores, ",")
	}

	servidorReal := grpc.NewServer()
	db.RegisterBaseServer(servidorReal, db.NuevoServidor(*addr, *esLider, ListaSeguidores))
	// TODO: completar para usar argumentos
	listen, err := net.Listen("tcp", *addr) 

	if err == nil {
		log.Printf("Servidor escuchando en puerto %s", *addr)
	}

	if err != nil {
		log.Fatalf("fallo al escuchar: %v", err)
	}
	
	if err := servidorReal.Serve(listen); err != nil {
		log.Fatalf("fallo al servir: %v", err)
	}
}

