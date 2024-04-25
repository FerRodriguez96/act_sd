package pkg

import (
	"context"

)

// La implementación del servidor
type Servidor struct {
	UnimplementedContadorServer
	// TODO: Defina un contador
	contador int32;
}

func NuevoServidor() Servidor {
	// TODO: Debe retonar una instancia del servidor definida previamente.
	// Complete de ser necesario.
	return Servidor{
		contador: 0,
	}
}

// Implementación de Obtener definido en el archivo `.proto`.
// TODO: Implementar `Obtener`. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s Servidor) Obtener(ctx context.Context, msg *Vacio) (*Valor, error) {

	retorno := Valor{
		contador: s.contador,
	}

	return &retorno, nil
}

// Implementación de Incrementar definido en el archivo `.proto`.
// TODO: Implementar 'Incrementar'. Si se produce algún error, devuelva el mensaje de error
// que desee.
func (s Servidor) Incrementar(ctx context.Context, _ *Vacio) (*Valor, error) {
	return nil, nil
}
