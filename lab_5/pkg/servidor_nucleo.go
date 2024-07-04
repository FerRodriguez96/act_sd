package pkg

import (
	"context"
	"errors"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// La implementación del servidor
type Servidor struct {
	UnimplementedBaseServer
	// TODO: Defina lo necesario para administrar los datos desde el servidor
	store map[string][]byte
	nodos []string
	mu sync.Mutex
}

func NuevoServidor(nodos []string) *Servidor {
	// TODO: Debe retonar una instancia del servidor definida previamente.
	for _, nodo := range nodos {
		conn, _ := grpc.Dial(
			nodo,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

	}
	// Complete de ser necesario.
	return &Servidor{
		store: make(map[string][]byte),
		nodos: nodos,
	}
}

// Implementación de put definido en el archivo `.proto`.
// TODO: Implementar `Put`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) Put (ctx context.Context, msg *ParametroPut) (*ResultadoPut, error) {
	//verificar si la clave es nula
	if msg.Clave == ""{
		return nil, errors.New("la clave no puede ser nula")
	}
	// verificar si el valor es nulo
	if msg.Valor == nil{
		return nil, errors.New("el valor no puede ser nulo")
	}
	// cerrojo para proteger la asignacion
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[msg.Clave] = msg.Valor

	return &ResultadoPut{Mensaje : "Valor guardado exitosamente"}, nil
}

// Implementación de get definido en el archivo `.proto`.
// TODO: Implementar `Get`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) Get(ctx context.Context, msg *ParametroGet) (*ResultadoGet, error) {
	Valor, exists := s.store[msg.Clave]
	if !exists {
		return nil, errors.New("no se encontro la clave")
	}
	return &ResultadoGet{Valor: Valor}, nil
}

// Implementación de getAll definido en el archivo `.proto`.
// TODO: Implementar `GetAll`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) GetAll(ctx context.Context, msg *ParametroGetAll) (*ResultadoGetAll, error) {
	if len(s.store) == 0 {
		return nil, errors.New("no existen valores almacenados")
	}
	return &ResultadoGetAll{Datos: s.store}, nil
}
