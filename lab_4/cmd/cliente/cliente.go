package main

import (
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
	// ingrese clave: 2, valor: Sist. Operativos
	// obtenga 2 e imprima lo devuelto
	// obtenga 3 e imprima lo devuelto
	// ingrese clave: 2, valor: Sistemas Operativos
	// obtenga todos los datos almacenados e imprima lo devuelto
}
