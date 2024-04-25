package main

import (
	"context"
	"fmt"
	mensajero "mensajero/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establece una conexión con el servidor
	conexion, _:= grpc.Dial(
		// dirección del servidor
		"localhost:8000",
		// indica que se debe conectar usando TCP sin SSL
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// bloquea el hilo hasta que la conexión se establezca
		grpc.WithBlock(),
	)
	// crea un nuevo cliente gRPC sobre la conexión
	cliente := mensajero.NewMensajeroClient(conexion)
	// envia un mensaje que retorna un TokenAutenticacion; se bloquea hasta que termina
	auth, _:= cliente.Conectar(
		// miren la biblioteca standard `context`
		context.Background(),
		// nuestros datos (pasa como contenido de UsuarioOrigen el String usuario)
		&mensajero.Registracion{UsuarioOrigen: "usuario"},
	)
	// debe imprimir: Token: "usuario se ha autenticado"
	fmt.Println(auth)
	}
