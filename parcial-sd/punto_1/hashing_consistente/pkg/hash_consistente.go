package pkg

import (
	"crypto/sha256"
	"encoding/binary"
)

var (
	// datos a almacenar
	datos map[uint32]string = make(map[uint32]string)
	// nodos del sistema 
	nodos map[uint32]string = make(map[uint32]string)
	clave_nodos []uint32 = make([]uint32, 0)
)

// genera el hash de un string usando sha256
func hash(s string) uint32 {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return binary.BigEndian.Uint32(hashed)
}

// agrega un nodo al sistema
func AgregarNodo(s string) {
	hash := hash(s)
	nodos[hash] = s
	clave_nodos = append(clave_nodos, hash)
}

// agrega un dato al sistema
func AgregarDato(s string) {
	hash := hash(s)
	datos[hash] = s
}

func GetSucesor(s string) string {
	// dado un string s, determine el nodo que le corresponde
	// usando hashing consistente
	// debe retornar el nodo que le corresponde

	hash := hash(s)
	for _, clave := range clave_nodos {
		if clave >= hash {
			return nodos[clave]
		}
	}
	return nodos[clave_nodos[0]]

}

func GetNodos() map[uint32]string {
	return nodos
}

func GetDatos() map[uint32]string {
	return datos
}
