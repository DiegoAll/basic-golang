package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	//Una forma de instanciar un struct
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Existe otra forma para instanciarlo como una clase vacia
	var otherCar car
	otherCar.brand = "Aston Martin"
	otherCar.year = 2021
	fmt.Println(otherCar)

}
