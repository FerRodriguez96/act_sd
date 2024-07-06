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
	direccionLideres []string
)

func main() {

	//definicion de array de direcciones
	direccionLideres = []string{"localhost:8000", "localhost:8001", "localhost:8002"}

	// insercion de claves aleatorias
	for i := 0; i < 1000; i++ {
		clave := generarClaveAleatoria()
		valor := "Valor" + strconv.Itoa(i)
		Put (clave, valor)
	}

	// Contar claves en cada nodo
	contarClaves()
}

//funcion hash
func claveHash(clave string) int {
	hash := sha256.New()
	hash.Write([]byte(clave))
	hashByte := hash.Sum(nil)
	hashInt := new(big.Int).SetBytes(hashByte)
	cant_serv := len(direccionLideres)
	
	if cant_serv == 0{
		log.Fatal("No hay servidores disponibles")
	}

	index := int(hashInt.Int64() % int64(cant_serv))

	if index < 0 {
		index = -index
	}
	return index
}
//asignacion de la clave hash a los nodos
func NodoParaClave(clave string) string {
	index := claveHash(clave)
	if index < 0 || index >= len(direccionLideres){
		log.Fatalf("Indice de servidor fuera de rango: %d", index)
	}
	return direccionLideres[index]
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
    cliente := db.NewBaseClient(conexion)

    _, err = cliente.Put(context.Background(), &db.ParametroPut{Clave: clave, Valor: []byte(valor)})
    if err != nil {
        log.Fatalf("Error al insertar clave %s en el nodo %s: %v", clave, nodo, err)
    }
    fmt.Printf("Clave %s insertada en el nodo %s\n", clave, nodo)
}

func contarClaves(){
	for _, nodo := range direccionLideres{
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
