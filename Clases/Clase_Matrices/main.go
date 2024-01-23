package main

import "fmt"

func main() {

	// intanciamos un vector
	Calificaciones := [5]int{10, 5, 2, 7, 9}

	for i, v := range Calificaciones {
		fmt.Printf("El alumno %d, tiene de calificacion %d \n", i+1, v)
	}

	//Vector indefinido

	//Arboles := [...]int{1, 2, 3, 4, 5}

	//fmt.Println(Arboles)

}
