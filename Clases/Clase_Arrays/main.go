package main

import "fmt"

func main() {
	/*
		var Calificaciones [5]int

		fmt.Println(Calificaciones)

		Calificaciones[0] = 10

		Calificaciones[5] = 2

		fmt.Println(Calificaciones)
	*/

	// slice
	// make([]tipo, )
	// Como primer parametro lo que yo quiero crear

	Productos := make([]string, 3)

	Productos[0] = "Sombrero"

	Productos[1] = "Zapatillas"

	Productos[2] = "Remera"

	fmt.Println(len(Productos))
	fmt.Println(Productos)

	Productos = append(Productos, "Cinturones")

	fmt.Println(len(Productos))
	fmt.Println(Productos)

	//fmt.Println(Productos)

}
