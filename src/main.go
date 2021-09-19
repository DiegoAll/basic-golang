package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println()

	/* Suponiendo que yo quiera modificar la RAM podria utilizar la instrucción  myPC.ram = 20
	Sin embargo a nivel de programación, a gran escala y teniendo mucho codigo. esto incluso puede llegar a ser redundante.
	Entonces se utiliza la siguiente técnica para optimizar. Crear una funcion como la siguiente: "duplicateRAM"
	*/
	fmt.Println("Actual ==>", myPC)

	myPC.duplicateRAM()
	fmt.Println("Después de la primera invocación ==>", myPC)

	myPC.duplicateRAM()
	fmt.Println("Después de la segunda invocación ==>", myPC)
}
