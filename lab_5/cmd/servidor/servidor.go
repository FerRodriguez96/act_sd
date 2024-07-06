package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	db "db/pkg"
)

type ConfiguracionServidor struct {
	Direccion string
	EsLider bool
	Seguidores []string
}

// Configuración de los nodos
var ConfiguracionServidores = []ConfiguracionServidor{
    {
        Direccion:  "localhost:8000",
        EsLider:    true,
        Seguidores: []string{"localhost:8003", "localhost:8004"},
    },
    {
        Direccion:  "localhost:8001",
        EsLider:    true,
        Seguidores: []string{"localhost:8005", "localhost:8006"},
    },
    {
        Direccion:  "localhost:8002",
        EsLider:    true,
        Seguidores: []string{"localhost:8007", "localhost:8008"},
    },
    {Direccion: "localhost:8003", EsLider: false, Seguidores: nil},
    {Direccion: "localhost:8004", EsLider: false, Seguidores: nil},
    {Direccion: "localhost:8005", EsLider: false, Seguidores: nil},
    {Direccion: "localhost:8006", EsLider: false, Seguidores: nil},
    {Direccion: "localhost:8007", EsLider: false, Seguidores: nil},
    {Direccion: "localhost:8008", EsLider: false, Seguidores: nil},
}

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

