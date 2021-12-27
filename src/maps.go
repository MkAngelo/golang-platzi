package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["Jose"] //Si no existe pone zeroValue
	fmt.Println(value, ok)

}
