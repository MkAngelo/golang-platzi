package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de variables enteras

	base := 2           //No ha sido declarada anteriormente creada e iniciada sin tipo de dato
	var altura int = 14 // declaramos la variable y asignamos un valor con el tipo de variable
	var area int        // solo la declaramos

	fmt.Println(base, altura, area)

	// Zero values
	var a int     // 0
	var b float64 // 0
	var c string  // ''
	var d bool    // false
	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

}
