package main

import (
	"context"
	"log"
	mensajero "mensajero/pkg"
	"net"
	"google.golang.org/grpc"
)

// un nuevo tipo que implementa mensajero.UnimplementedMensajeroServer
type servidor struct {
	mensajero.UnimplementedMensajeroServer
}

// un ejemplo de implementación de la interfaz del servidor
func (s servidor) Conectar(_ context.Context, r *mensajero.Registracion) (*mensajero.TokenAutenticacion, error) {
	token := r.UsuarioOrigen + " se ha autenticado"
	return &mensajero.TokenAutenticacion{Token: token}, nil
}

func main() {
	servidorReal := grpc.NewServer()
	mensajero.RegisterMensajeroServer(servidorReal, servidor{})
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("fallo al escuchar: %v", err)
	}
	if err := servidorReal.Serve(listen); err != nil {
		log.Fatalf("fallo al servir: %v", err)
	}
}