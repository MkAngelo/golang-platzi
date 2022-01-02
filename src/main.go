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

func (myPC pc) show() {
	showString := fmt.Sprintf("Esta PC es de la marca %s, tiene %d GB de RAM y %d GB de disco.", myPC.brand, myPC.ram, myPC.disk)
	fmt.Println(showString)
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

	myPC.show()
	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
}
