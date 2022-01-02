package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong.")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a // Puntero a memoria

	fmt.Println(a)
	fmt.Println(*b) // Accedemos al valor en memoria
	fmt.Println(b)  // Accedemos al puntero

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "HP"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
}
