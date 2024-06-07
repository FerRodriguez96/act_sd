package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	db "db/pkg"
)

const (
	DIRECCION_SERVIDOR_PREDETERMINADA string = "localhost"
	PUERTO_SERVIDOR_PREDETERMINADO    string = "8000"
)

var (
	puertoServidor    string
	direccionServidor string
)

func main() {
	flag.StringVar(&puertoServidor, "p", PUERTO_SERVIDOR_PREDETERMINADO, "puerto a conectarse")
	flag.StringVar(&direccionServidor, "d", DIRECCION_SERVIDOR_PREDETERMINADA, "dirección del servidor")
	flag.Parse()

	// Validar que el puerto se encuentre entre 1 y 65535
	if !validarPuerto(puertoServidor) {
		log.Fatalf("El puerto %s no está en el rango de 1 a 65535", puertoServidor)
	}

	direccion := fmt.Sprintf("%s:%s", direccionServidor, puertoServidor)

	// Establece una conexión con el servidor
	conexion, _ := grpc.Dial(
		// dirección del servidor
		direccion,
		// indica que se debe conectar usando TCP sin SSL
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// bloquea el hilo hasta que la conexion se establezca
		grpc.WithBlock(),
	)

	// Crea un nuevo cliente gRPC sobre la conexión
	cliente := db.NewBaseClient(conexion)

	// Ingrese clave: 1, valor: Sistemas Distribuidos
	cliente.Put(context.Background(), &db.ParametroPut{Clave: "1", Valor: []byte("Sistemas Distribuidos")})

	// Ingrese clave: 2, valor: Sist. Operativos
	cliente.Put(context.Background(), &db.ParametroPut{Clave: "2", Valor: []byte("Sist. Operativos")})

	// Obtenga clave: 2 e imprima lo devuelto
	resultadoGet, err := cliente.Get(context.Background(), &db.ParametroGet{Clave: "2"})
	if err != nil {
		log.Printf("Error al obtener clave 2: %v", err)
	} else {
		fmt.Printf("Valor para clave 2: %s\n", string(resultadoGet.Valor))
	}

	// Obtenga clave: 3 e imprima lo devuelto
	resultadoGet, err = cliente.Get(context.Background(), &db.ParametroGet{Clave: "3"})
	if err != nil {
		log.Printf("Error al obtener clave 3: %v", err)
	} else {
		fmt.Printf("Valor para clave 3: %s\n", string(resultadoGet.Valor))
	}

	// Ingrese clave: 2, valor: Sistemas Operativos
	_, err = cliente.Put(context.Background(), &db.ParametroPut{Clave: "2", Valor: []byte("Sistemas Operativos")})
	if err != nil {
		log.Fatalf("Error al actualizar clave 2: %v", err)
	}

	// Obtenga todos los datos almacenados e imprima lo devuelto
	resultadoGetAll, err := cliente.GetAll(context.Background(), &db.ParametroGetAll{})
	if err != nil {
		log.Fatalf("Error al obtener todos los datos: %v", err)
	} else {
		fmt.Println("Datos almacenados:")
		for clave, valor := range resultadoGetAll.Datos {
			fmt.Printf("Clave: %s, Valor: %s\n", clave, string(valor))
		}
	}
}

func validarPuerto(puerto string) bool {
	p, err := strconv.Atoi(puerto)
	if err != nil {
		return false
	}
	return p > 0 && p <= 65535
}
