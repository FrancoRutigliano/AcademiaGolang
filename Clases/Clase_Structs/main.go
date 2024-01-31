package main

import "fmt"

// siempre se suele declara fuera de una funcion
// buena practica
type Persona struct {
	Nombre string
	Edad   int
	Peso   int
	Altura int
	Genero string
}

// metodos
// receiver
func (p Persona) saludar(nombre string) {
	fmt.Printf("Hola %s, te saluda %s", nombre, p.Nombre)
}

func (p Persona) CambiarEdad() {
	p.Edad = 21
}

func main() {
	var persona1 Persona

	persona1.Nombre = "Franco"

	persona1.Edad = 20

	nombre := "Mati"

	persona1.saludar(nombre)

	persona1.CambiarEdad()

	fmt.Println("Mi edad ", persona1.Edad)

}
