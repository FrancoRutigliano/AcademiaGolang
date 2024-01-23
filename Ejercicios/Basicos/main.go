package main

import "fmt"

// una manera
const (
	TiempoHorno = 40 // camel case
)

func main() {

	numeroCapas := 3

	transcurrido := 20

	TiempoTotalDeCapa := TiempoDePreparacion(numeroCapas)

	tiempoRestante := TiempoRestante(transcurrido)

	TiempoTotal := TiempoTranscurrido(numeroCapas, transcurrido)

	fmt.Println("TIEMPO TOTAL DE CAPAS: ->>", TiempoTotalDeCapa)

	fmt.Println("TIEMPO RESTANTE  ->>", tiempoRestante)

	fmt.Println("TIEMPO TOTAL EN EL HORNO: ->> ", TiempoTotal)
}

// cada capa 2 minutos
func TiempoDePreparacion(numeroDeCapas int) int {
	TiempoDeCapa := 2

	totalTiempo := numeroDeCapas * TiempoDeCapa

	return totalTiempo
}

func TiempoRestante(tiempo int) int {
	tiempoRestante := TiempoHorno - tiempo
	// validacion ----> por si el numero llega a ser negativo

	// ----> no funciona
	if tiempoRestante < 0 {
		return 0
	}
	return tiempoRestante

}

func TiempoTranscurrido(numeroDeCapas, tiempoActual int) int {
	tiempoCapa := TiempoDePreparacion(numeroDeCapas)
	totalTiempo := tiempoCapa + tiempoActual

	return totalTiempo
}
