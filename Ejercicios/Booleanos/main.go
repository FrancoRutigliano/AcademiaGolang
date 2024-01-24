package main

import (
	"fmt"
	"log"
	"strings"

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

// ------ Hacer Tolower -------
// strings.Tolower()
var (
	Estado string
)

func main() {

	fmt.Print("Estado del Caballero: (si o no) ")
	fmt.Scan(&Estado)
	Estado = strings.ToLower(Estado)
	if Estado == "si" {
		// caballero esta despierto
		caballeroEstado = true
		fmt.Println("El caballero esta despierto")
	} else if Estado == "no" {
		caballeroEstado = false
		fmt.Println("El caballero esta dormido")
	} else {
		log.Fatal("Estado incorrecto, intente nuevamente")
	}

	fmt.Print("Estado del Aquero: (si o no) ")
	fmt.Scan(&Estado)
	Estado = strings.ToLower(Estado)

	if Estado == "si" {
		// arquero esta despierto
		arqueroEstado = true
		fmt.Println("El arquero esta despierto")
	} else if Estado == "no" {
		arqueroEstado = false
		fmt.Println("El arquero esta dormido")
	} else {
		log.Fatal("Estado incorrecto, intente nuevamente")
	}

	fmt.Println("!ES MOMENTO DE TOMAR UNA DECISION!")
	// del menu
	var opc int = 0
	fmt.Print("Selecionar una opcion: 1) Ataque rapido | 2) Espiar | 3) Salvar al prisionero \n")
	fmt.Scan(&opc)
	switch opc {
	case 1:
		funciones.AtaqueRapido(caballeroEstado)
	case 2:
		funciones.Espiar(caballeroEstado, arqueroEstado)
	case 3:
		fmt.Print("Esta perro presente: (si o no): ")
		fmt.Scan(&Estado)
		Estado = strings.ToLower(Estado)

		if Estado == "si" {
			// caballero esta despierto
			perroEstado = true
		} else if Estado == "no" {
			perroEstado = false
		} else {
			log.Fatal("Estado incorrecto, intente nuevamente")
		}

		fmt.Print("\nEstado del prisionero: (si o no):")
		fmt.Scan(&Estado)
		Estado = strings.ToLower(Estado)

		if Estado == "si" {
			// caballero esta despierto
			priosioneroEstado = true
		} else if Estado == "no" {
			priosioneroEstado = false
		} else {
			log.Fatal("Estado incorrecto, intente nuevamente")
		}

		funciones.SalvarPrisionero(caballeroEstado, arqueroEstado, perroEstado, priosioneroEstado)

	default:
		fmt.Println("Número incorrecto, intente nuevamente")
	}

}

// func corroborarEstado(nombre string) bool {
// 	fmt.Printf("Estado del %s: (si o no) ", nombre)
// 	fmt.Scan(&Estado)

// 	Estado = strings.ToLower(Estado)
