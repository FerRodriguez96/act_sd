package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	db "db/pkg"
)

const (
	PUERTO_SERVIDOR_PREDETERMINADO    string = "8000"
)

var (
	puertoServidor    string
	direccionServidores []string
)

func main() {
	direccionServidores = []string{"localhost:8000", "localhost:8001", "localhost:8002"}

	// insercion de claves aleatorias
	for i := 0; i < 1000; i++ {
		clave := generarClaveAleatoria()
		valor := "Valor" + strconv.Itoa(i)
		Put (clave, valor)
	}

	// Contar claves en cada nodo
	contarClaves()
}

func claveHash(clave string) int {
	hash := sha256.New()
	hash.Write([]byte(clave))
	hashByte := hash.Sum(nil)
	hashInt := new(big.Int).SetBytes(hashByte)
	return int(hashInt.Int64() % int64(len(direccionServidores)))
}

func NodoParaClave(clave string) string {
	index := claveHash(clave)
	return direccionServidores[index]
}

func Put(clave, valor string) {
    nodo := NodoParaClave(clave)
    conexion, err := grpc.Dial(
        nodo,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock(),
    )
    if err != nil {
        log.Fatalf("No se pudo conectar al servidor en %s: %v", nodo, err)
    }

    defer conexion.Close()
    client := db.NewBaseClient(conexion)

    _, err = client.Put(context.Background(), &db.ParametroPut{Clave: clave, Valor: []byte(valor)})
    if err != nil {
        log.Fatalf("Error al insertar clave %s en el nodo %s: %v", clave, nodo, err)
    }
    fmt.Printf("Clave %s insertada en el nodo %s\n", clave, nodo)
}

func validarPuerto(puerto string) bool {
	p, err := strconv.Atoi(puerto)
	if err != nil {
		return false
	}
	return p > 0 && p <= 65535
}

func contarClaves(){
	for _, nodo := range direccionServidores {
		conexion, err := grpc.Dial(
			nodo,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
	
		if err != nil {
			log.Fatalf("No se pudo conectar al servidor en %s: %v", nodo, err)
		}

		defer conexion.Close()
		cliente := db.NewBaseClient(conexion)

		resultadoGetAll, err := cliente.GetAll(context.Background(), &db.ParametroGetAll{}) 
		if err != nil {
			log.Fatalf("Error al obtener todos los datos del nodo %s: %v", nodo, err)
		} else {
			fmt.Printf("Nodo %s tiene %d claves almacenadas\n", nodo, len(resultadoGetAll.Datos))
		}

	}	
}

func generarClaveAleatoria() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%x", rand.Int63())
}
