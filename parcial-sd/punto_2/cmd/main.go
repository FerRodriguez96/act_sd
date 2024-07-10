package main

import (
	"fmt"
)

func CompararRelojes(reloj1, reloj2 map[string]int) int {
	igual := true
	menor := true
	mayor := true

	for proceso, tiempo1 := range reloj1 {
		tiempo2, existe := reloj2[proceso]
		if !existe {
			tiempo2 = 0
		}

		if tiempo1 != tiempo2 {
			igual = false
		}
		if tiempo1 > tiempo2 {
			menor = false
		}
		if tiempo1 < tiempo2 {
			mayor = false
		}
	}

	if igual {
		return 0
	}
	if menor && !mayor {
		return 1
	}
	if mayor && !menor {
		return 2
	}
	return -1
}

func main() {
	reloj1 := map[string]int{
		"P1": 1,
		"P2": 3,
		"P3": 2,
		"P4": 4,
		"P5": 5,
	}

	reloj2 := map[string]int{
		"P1": 2,
		"P2": 3,
		"P3": 2,
		"P4": 4,
		"P5": 5,
	}

	resultado := CompararRelojes(reloj1, reloj2)
	fmt.Println("Resultado de la comparación:", resultado)
}