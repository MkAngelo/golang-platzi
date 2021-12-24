package main

import "fmt"

func isPair(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

func validatePass(pass, confPass string) string {
	message := "No coinciden"
	if pass == confPass {
		message = "Coinciden"
	}
	return message
}

func main() {
	result := isPair(5)
	fmt.Println(result)
	password := validatePass("Hola", "hola")
	fmt.Println(password)
}
