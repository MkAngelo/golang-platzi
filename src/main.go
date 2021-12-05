package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de variables enteras

	base := 2           //creada e iniciada sin tipo de dato
	var altura int = 14 //creada e iniciada con valor
	var area int        // creada

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

}
