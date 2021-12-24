package main

import "fmt"

func main() {
	valor1 := 2
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("El valor es 1")
	} else {
		fmt.Println("El valor NO es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es Verdad")
	} else {
		fmt.Println("Es Mentira")
	}

	// With and
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Es Verdad")
	} else {
		fmt.Println("Es Mentira")
	}
}
