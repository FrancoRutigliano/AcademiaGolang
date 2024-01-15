package main

import "fmt"

func main() {
	velocidadProduccion := 1547
	tasaExito := 90
	cantidadDeAutos := 2

	AutosExitosos := AutosPorHora(velocidadProduccion, tasaExito)

	AutosExitososPorMin := AutosPorMinutos(1105, 90)

	costosAutos := CalcularCosto(cantidadDeAutos)

	fmt.Println(AutosExitosos)

	fmt.Println("Autos exitosos por minuto -->", AutosExitososPorMin)

	fmt.Printf("CANTIDAD DE AUTOS %d ---- COSTOS $%d", cantidadDeAutos, costosAutos)

}

func AutosPorHora(velocidadProduccion, tasaExito int) float32 {
	autosExitosos := velocidadProduccion * tasaExito / 100

	return float32(autosExitosos)
}

// tomar el num de autos producidos x hora
// tasa de exito
// por hora cuantos autos exitosos se producen (int)

func AutosPorMinutos(velocidadProduccion, tasaExito int) int {
	AutosPorMin := velocidadProduccion / 60

	autosExitosos := tasaExito * AutosPorMin / 100

	return autosExitosos
}

// Cada auto cuesta 10.000 dolares si tiene exito cuesta lo mismo que si no.
// 1 pack de autos  = 10 autos salen 95

func CalcularCosto(cantAutos int) int {
	// cantidad de autos / por grupos
	// 37 me devuelve 3
	pack, ind := 95000, 10000
	var individuales int = 0
	if cantAutos < 10 {
		individuales = cantAutos % 10
		return individuales * ind
	} else {
		grupos := cantAutos / 10
		individuales = cantAutos % 10
		return (grupos * pack) + (individuales * ind)
	}

}
