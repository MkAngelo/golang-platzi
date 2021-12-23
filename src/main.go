package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int { // Valor de salida
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "H0l4!")

	value := returnValue(10)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(5)
	fmt.Printf("Value1: %d, Value2: %d\n", value1, value2)

	//PAra evitar usar una variable de retorno con '_'
}
