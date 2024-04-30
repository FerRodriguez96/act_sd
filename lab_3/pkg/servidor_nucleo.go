package pkg

import (
	"context"
)

// La implementación del servidor
type Servidor struct {
	UnimplementedContadorServer
	// TODO: Defina un contador
	Contador int32;
}

func NuevoServidor() *Servidor {
	// TODO: Debe retonar una instancia del servidor definida previamente.
	// Complete de ser necesario.
	return &Servidor{
		Contador: 0,
	}
}

// Implementación de Obtener definido en el archivo `.proto`.
// TODO: Implementar `Obtener`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s Servidor) Obtener(ctx context.Context, msg *Vacio) (*Valor, error) {

	valor := Valor{
		Contador: s.Contador,
	}

	return &valor, nil
}

// Implementación de Incrementar definido en el archivo `.proto`.
// TODO: Implementar 'Incrementar'. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s Servidor) Incrementar(ctx context.Context, _ *Vacio) (*Valor, error) {

	s.Contador++

	retorno := Valor{
		Contador : s.Contador,
	}

	return &retorno, nil
}
