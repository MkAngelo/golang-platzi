package main

import "fmt"

func main() {
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	// Division
	result = y / x
	fmt.Println("Division: ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Aumenta
	x++
	fmt.Println("Aumenta: ", x)

	// Disminiye
	x--
	fmt.Println("Disminuye: ", x)
}
