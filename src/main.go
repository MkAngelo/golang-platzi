package main

import "fmt"

func main() {
	switch modulo := 10 % 2; modulo {
	case 0:
		fmt.Println("PAR")
	default:
		fmt.Println("IMPAR")
	}

	// Sin condiciones

	value := 101
	switch {
	case value == 100:
		fmt.Println("El valor es 100")
	case value > 100:
		fmt.Println("El valor es mayor a 100")
	default:
		fmt.Println("El valor es menor a 100")
	}
}
