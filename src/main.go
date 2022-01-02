package main

import (
	"curso_golang_platzi/src/mypackage"
	"fmt"
	// pk "ruta" es la forma de poner alias
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2001
	// Solo se  pude acceder a los paquetes publicos

	fmt.Println(myCar)
}
