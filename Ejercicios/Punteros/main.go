package main

import "fmt"

type Persona struct {
	Nombre   string
	Edad     int
	Apellido string
}

type Alumno struct {
	Persona
	Legajo_id int
	materias  []string
}

type Alumnos []Alumno

func (A Alumno) obtenerNombreCompleto() {
	fmt.Println(A.Persona.Nombre, A.Persona.Apellido)
}

//EJERCICIO
func (A Alumno) obtenerMaterias() {

}

func (A Alumno) obtenerLegajo() {

}

func main() {
	alumno := Alumno{
		Persona: Persona{
			Nombre:   "Franco",
			Edad:     21,
			Apellido: "Rutigliano",
		},
		Legajo_id: 10,
		materias:  []string{"Matemáticas", "Historia", "Programación"},
	}

	alumno2 := Alumno{
		Persona: Persona{
			Nombre:   "Matias",
			Edad:     20,
			Apellido: "Lloret",
		},
		Legajo_id: 11,
		materias:  []string{"Matemáticas", "Historia", "Programación"},
	}

	fmt.Println(alumno.Persona.Nombre)

	fmt.Println(alumno2.Persona.Nombre)

	alumno.obtenerNombreCompleto()

	alumno2.obtenerNombreCompleto()

}
