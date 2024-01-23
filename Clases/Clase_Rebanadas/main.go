package main

import (
	"fmt"
	"strings"
)

//diasSemana := []string{"Domingo", "Sabado", "Lunes", "Martes"}

// fmt.Println(diasSemana)

// diasSemana = append(diasSemana, "Miercoles", "Jueves", "Viernes")

// fmt.Println(diasSemana)

// diasSemana = append(diasSemana, "otro día")

// fmt.Println(diasSemana)

// diasSemana = append(diasSemana[:5], diasSemana[6:]...)

// fmt.Println(diasSemana)

func main() {
	var opc int = 0
	var equipo int = 0
	var des, nombreEquipo string
	equiposDeFutbol := []string{
		"Argentinos Juniors",
		"Atlético Tucumán",
		"Banfield",
		"Barracas Central",
		"Belgrano (Córdoba)",
		"Boca Juniors",
		"Central Córdoba (Santiago del Estero)",
		"Defensa y Justicia",
		"Deportivo Riestra",
		"Estudiantes de La Plata",
		"Gimnasia La Plata",
		"Godoy Cruz Antonio Tomba",
		"Huracán",
		"Independiente",
		"Independiente Rivadavia",
		"Instituto (Córdoba)",
		"Lanús",
		"Newell's Old Boys",
		"Platense",
		"Racing Club",
		"River Plate",
		"Rosario Central",
		"San Lorenzo",
		"Sarmiento (Junín)",
		"Talleres (Córdoba)",
		"Tigre",
		"Unión (Santa Fe)",
		"Vélez Sarsfield",
	}

	for {
		fmt.Println("------------Menu------------")
		fmt.Printf("1)Eliminar un equipo 1 al %v\n", len(equiposDeFutbol))
		fmt.Printf("2)Agregar un equipo 1 al %v\n", len(equiposDeFutbol))
		fmt.Println("3) Mostrar todos los equipos ")
		fmt.Println("4) Salir")
		fmt.Print("Elegir una acción: ")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			fmt.Print("\nIngrese el número de equipo a eliminar: ")
			fmt.Scan(&equipo)

			for i, v := range equiposDeFutbol {
				if i == equipo-1 {
					fmt.Println("Equipo a eliminar: ", v)
				}
			}

			fmt.Print("\nEstas seguro: (si o no) ")
			fmt.Scan(&des)

			if strings.ToLower(des) == "si" {
				equiposDeFutbol = append(equiposDeFutbol[:equipo-1], equiposDeFutbol[equipo:]...)
			} else if strings.ToLower(des) == "no" {
				fmt.Println("No se eliminará")
			}

		case 2:

			fmt.Print("Que equipo queres agregar: ")
			fmt.Scan(&nombreEquipo)

			fmt.Print("\nEstas seguro: (si o no) ")
			fmt.Scan(&des)

			if strings.ToLower(des) == "si" {
				equiposDeFutbol = append(equiposDeFutbol, nombreEquipo)
			} else if strings.ToLower(des) == "no" {
				fmt.Println("Elija otra opcion")
			}

		case 3:
			fmt.Print("\n")
			for _, valor := range equiposDeFutbol {
				fmt.Println(valor)
			}
			fmt.Print("\n")
		case 4:
			return
		default:
			fmt.Println("Numero incorrecto")

		}
	}
}
