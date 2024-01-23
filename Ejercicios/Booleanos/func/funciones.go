package funciones

import "fmt"

func AtaqueRapido(caballeroEstado bool) {
	if caballeroEstado {
		fmt.Println("Puedes hacer un ataque rapido")
	} else {
		fmt.Println("No podemos hacer un ataque rapido, el caballero esta despierto")
	}
}
