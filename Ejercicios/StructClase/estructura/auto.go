package estructura

import (
	"fmt"
	"strings"
)

type Auto struct {
	Marca           string
	Color           string
	Modelo          string
	Año             int
	VelocidadMaxima int
	Vendido         bool
}

func (a Auto) ObtenerDatos() {
	fmt.Printf("Marca: %s\nModelo: %s\nColor: %s\nAño: %d\nVelocidad Maxima: %d", a.Marca, a.Modelo, a.Color, a.Año, a.VelocidadMaxima)
}

type Deportivo struct {
	Auto
	Aleron bool // true es tiene false es no tiene
}

type Deportivos []Deportivo

func (d *Deportivos) crearDeportivos() {

	var marca, modelo, color, aleron string
	var año, velocidadMaxima int
	var aleronBool bool

	fmt.Println("MARCA: ")
	fmt.Scanln(&marca)
	fmt.Println("MODELO:")
	fmt.Scanln(&modelo)
	fmt.Println("COLOR: ")
	fmt.Scanln(&color)
	fmt.Println("AÑO: ")
	fmt.Scanln(&año)
	fmt.Println("Velocidad maxima: ")
	fmt.Scanln(&velocidadMaxima)
	fmt.Println("Aleron: ")
	fmt.Scanln(&aleron)
	aleron = strings.ToLower(aleron)

	if aleron == "si" {
		aleronBool = true
	} else {
		aleronBool = false
	}

	deportivo := Deportivo{
		Auto: Auto{
			Marca:           marca,
			Modelo:          modelo,
			Color:           color,
			Año:             año,
			VelocidadMaxima: velocidadMaxima,
		},
		Aleron: aleronBool,
	}

	*d = append(*d, deportivo)

}

func (d Deportivos) MostrarDeportivos() {
	// debería mostrarme todos los deportivos
}

func (d Deportivos) MostrarDeportivosVendidos() {
	// tenes que MOSTRAR los deportivos que tengan el campo deportivo vendido en true
}

func (d *Deportivos) DeportivoVendido(index int) {
	// MODIFICAR un deportivo al campo vendido en true
	// como se cual deportivo modificar? Por el index.
}

func (d *Deportivos) EliminarDeportivo(index int) {
	// eliminar un deportivo de la lista deportivos
	// como se cual deportivo eliminar? Por el index.
}

type Camioneta struct {
	Auto
	Traccion bool // true es 4*4 y false no lo es
}

type Camionetas []Camioneta

func (d *Camionetas) crearCamioneta() {

}
