package main

import (
	"fmt"
	"log"

	funciones "github.com/FrancoRutigliano/AcademiaGolang/Ejercicios/Booleanos/func"
)

// TAREA MATI:
// MEJORAR EL MENU.
// PODRÍAS CREAR FUNCIONES QUE CORROBOREN EL ESTADO DE CADA UNO DE LOS PERSONAJES

//----> opcional ---> modular al 100% el código

// si estan o no // si estan durmiendo o no
var (
	caballeroEstado   bool
	arqueroEstado     bool
	priosioneroEstado bool
	perroEstado       bool
)

var (
	Estado string
)

func main() {

	fmt.Print("Estado del Caballero: (si o no) ")
	fmt.Scan(&Estado)

	if Estado == "si" {
		// caballero esta despierto
		caballeroEstado = true
	} else if Estado == "no" {
		caballeroEstado = false
	} else {
		log.Fatal("Estado incorrecto")
	}

	fmt.Print("Estado del Aquero: (si o no) ")
	fmt.Scan(&Estado)

	if Estado == "si" {
		// arquero esta despierto
		arqueroEstado = true
	} else if Estado == "no" {
		arqueroEstado = false
	} else {
		log.Fatal("Estado incorrecto")
	}

	// del menu
	var opc int = 0
	fmt.Print("Selecionar opcion: 1) Ataque rapido | 2) Espiar | 3) Salvar al prisionero \n")
	fmt.Scan(&opc)
	switch opc {
	case 1:
		funciones.AtaqueRapido(caballeroEstado)
	case 2:
		Espiar(caballeroEstado, arqueroEstado)
	case 3:
		fmt.Print("Esta perro presente: (si o no): ")
		fmt.Scan(&Estado)

		if Estado == "si" {
			// caballero esta despierto
			perroEstado = true
		} else if Estado == "no" {
			perroEstado = false
		} else {
			log.Fatal("Estado incorrecto")
		}

		fmt.Print("\nEstado del prisionero: (si o no):")
		fmt.Scan(&Estado)

		if Estado == "si" {
			// caballero esta despierto
			priosioneroEstado = true
		} else if Estado == "no" {
			priosioneroEstado = false
		} else {
			log.Fatal("Estado incorrecto")
		}

		SalvarPrisionero(caballeroEstado, arqueroEstado, perroEstado, priosioneroEstado)

	default:
		fmt.Println("Número incorrecto")
	}

}

func Espiar(caballeroEstado, ArqueroEstado bool) {
	if caballeroEstado || arqueroEstado {
		fmt.Println("Podemos espiar")
	} else if !caballeroEstado && !arqueroEstado {
		fmt.Println("No te conviene espiar, los dos estan dormidos, preferible salvar al princionero")
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
		} else {
			fmt.Println("Se durmieron, volve mañana")
		}
	}
}

// func corroborarEstado(nombre string) bool {
// 	fmt.Printf("Estado del %s: (si o no) ", nombre)
// 	fmt.Scan(&Estado)

// 	Estado = strings.ToLower(Estado)
