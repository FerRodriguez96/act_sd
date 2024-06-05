package main

import (
	"context"
	db "db/pkg"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DIRECCION_SERVIDOR_PREDETERMINADA string = "localhost"
	PUERTO_SERVIDOR_PREDETERMINADO    string = "12345"
)

var (
	puertoServidor    string
	direccionServidor string
)

func main() {
	flag.StringVar(&puertoServidor, "p", PUERTO_SERVIDOR_PREDETERMINADO, "puerto a conectarse")
	flag.StringVar(&direccionServidor, "d", DIRECCION_SERVIDOR_PREDETERMINADA, "dirección del servidor")
	flag.Parse()

	// TODO: Validar que el puerto se encuentre entre 1 y 65535
	// De no estar en ese rango finalizar el programa

	direccion := fmt.Sprintf("%s:%s", direccionServidor, puertoServidor)

	// Establece una conexión con el servidor
	conexion, _ := grpc.Dial(
		// dirección del servidor
		direccion,
		// indica que se debe conectar usando TCP sin SSL
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// bloquea el hilo hasta que la conexión se establezca
		grpc.WithBlock(),
	)
	// crea un nuevo cliente gRPC sobre la conexión
	cliente := db.NewBaseClient(conexion)

	// TODO:
	// ingrese clave: 1, valor: Sistemas Distribuidos
	cliente.Put(context.Background(), &db.ParametroPut{Clave: "1", Valor: []byte("Sistemas Distribuidos")})
	// ingrese clave: 2, valor: Sist. Operativos
	cliente.Put(context.Background(), &db.ParametroPut{Clave: "2", Valor: []byte("Sist. Operativos")})
	// obtenga 2 e imprima lo devuelto
	cliente.Get(context.Background(), &db.ParametroGet{Clave: "2"})
	// obtenga 3 e imprima lo devuelto
	cliente.Get(context.Background(), &db.ParametroGet{Clave: "3"})
	// ingrese clave: 2, valor: Sistemas Operativos
	cliente.Put(context.Background(), &db.ParametroPut{Clave: "1", Valor: []byte("Sistemas Operativos")})
	// obtenga todos los datos almacenados e imprima lo devuelto
	cliente.GetAll(context.Background(), &db.ParametroGetAll{})
}
