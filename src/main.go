package main

import (
	"curso_golang_platzi/src/mypc"
	"fmt"
)

func newPC() {
	var myPC mypc.PcSettings
	myPC.Brand = "Lenovo"
	myPC.RAM = 12
	myPC.Disk = 500
	myPC.Year = 2020

	fmt.Println(myPC)
}

func main() {
	newPC()
}
