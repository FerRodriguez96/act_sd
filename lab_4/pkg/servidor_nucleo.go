package pkg

import (
	"context"
	"sync"
)

// La implementación del servidor
type Servidor struct {
	UnimplementedBaseServer
	// TODO: Defina lo necesario para administrar los datos desde el servidor
}

func NuevoServidor() Servidor {
	// TODO: Debe retonar una instancia del servidor definida previamente.
	// Complete de ser necesario.
	return Servidor{
	}
}

// Implementación de put definido en el archivo `.proto`.
// TODO: Implementar `Put`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) Put (ctx context.Context, msg *ParametroPut) (*ResultadoPut, error) {
}

// Implementación de get definido en el archivo `.proto`.
// TODO: Implementar `Get`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) Get(ctx context.Context, msg *ParametroGet) (*ResultadoGet, error) {
}

// Implementación de getAll definido en el archivo `.proto`.
// TODO: Implementar `GetAll`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) GetAll(ctx context.Context, msg *ParametroGetAll) (*ResultadoGetAll, error) {
}
