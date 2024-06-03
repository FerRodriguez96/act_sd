package pkg

import (
	"context"
)

// La implementación del servidor
type Servidor struct {
	UnimplementedBaseServer
	// TODO: Defina lo necesario para administrar los datos desde el servidor
	store map[string][]byte
}

func NuevoServidor() *Servidor {
	// TODO: Debe retonar una instancia del servidor definida previamente.
	// Complete de ser necesario.
	return &Servidor{
		store: make(map[string][]byte),
	}
}

// Implementación de put definido en el archivo `.proto`.
// TODO: Implementar `Put`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) Put (ctx context.Context, msg *ParametroPut) (*ResultadoPut, error) {
	s.store[msg.Clave] = []byte(msg.Valor)
	return &ResultadoPut{Mensaje : "Valor guardado exitosamente"}, nil
}

// Implementación de get definido en el archivo `.proto`.
// TODO: Implementar `Get`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) Get(ctx context.Context, msg *ParametroGet) (*ResultadoGet, error) {
	Valor, exists := s.store[msg.Clave]
	if !exists {
		return &ResultadoGet{Valor : nil}, nil
	}
	return &ResultadoGet{Valor: Valor}, nil
}

// Implementación de getAll definido en el archivo `.proto`.
// TODO: Implementar `GetAll`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s *Servidor) GetAll(ctx context.Context, msg *ParametroGetAll) (*ResultadoGetAll, error) {
	return &ResultadoGetAll{Datos: s.store}, nil
}
