package funciones

import "fmt"

func AtaqueRapido(caballeroEstado bool) {
	if !caballeroEstado {
		fmt.Println("Puedes hacer un ataque rapido")
	} else {
		fmt.Println("No podemos hacer un ataque rapido, el caballero esta despierto")
	}
}

func Espiar(caballeroEstado, ArqueroEstado bool) {
	if caballeroEstado && ArqueroEstado {
		fmt.Println("No te conviene espiar, los dos estan dormidos, preferible salvar al princionero")
	} else if caballeroEstado || !ArqueroEstado {
		fmt.Println("Podemos espiar")
	} else if !caballeroEstado || ArqueroEstado {
		fmt.Println("Podemos espiar")
	}
}

func SalvarPrisionero(caballeroEstado, arqueroEstado, perroEstado, prisioneroEstado bool) {
	// PRIMERO ---> ESTADO DEL PERRO
	// si el perro esta, para poder rescatar al prisionero, el arquero tiene que estar dormido
	if perroEstado {
		if !arqueroEstado && prisioneroEstado {
			fmt.Println("Salvalooooo")
		} else if arqueroEstado {
			fmt.Println("No podes, arquero despierto")
		} else if !prisioneroEstado {
			fmt.Println("Prisionero dormido")
		}
	} else {
		if prisioneroEstado && !caballeroEstado && !arqueroEstado {
			fmt.Println("Salvalooooo")
		} else if caballeroEstado && arqueroEstado {
			fmt.Println("Estan despiertos, no puedes salvarlo")
		} else {
			fmt.Println("Se durmieron, volve mañana")
		}
	}
}
