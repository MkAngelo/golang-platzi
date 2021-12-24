package main

import "fmt"

func main() {
	// Array (Inmutable)
	var array [4]int
	fmt.Println(array)

	for i := 0; i < 4; i++ {
		array[i] = i + 1
	}

	fmt.Println(array)

	// len -> elementos
	// cap -> capacidad maxima

	// Slices (Mutables)
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	// Metodos en el slice
	fmt.Println(slice[0]) // no es incluyente
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
