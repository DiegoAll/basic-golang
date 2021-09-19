package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

//Función para personalizar el output de consola.
/*func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}*/

func main() {
	myPC := pc{ram: 16, brand: "msi", disk: 100}
	fmt.Println(myPC)
}
