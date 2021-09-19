package main

import (
	pk "basic-golang/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Chevrolet"
	myCar.Year = 2021

	fmt.Println(myCar)

	pk.PrintMessage("Oelo")
}
