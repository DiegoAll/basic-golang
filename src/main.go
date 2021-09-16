package main

import "fmt"

func igualdad(v1, v2 bool) bool {
	return v1 == v2
}

func distinto(v1, v2 bool) bool {
	return v1 != v2
}

func comparacion(i1, i2 int) (ans1, ans2, ans3, ans4 bool) {
	return i1 > i2, i1 < i2, i1 >= i2, i1 <= i2
}

func logicGates(v1, v2 bool) (ans1, ans2, ans3 bool) {
	return v1 && v2, v1 || v2, !v1
}

func main() {

	fmt.Println("Operadores lógicos y de comparación")

	var valor1 bool = false
	var valor2 bool = true
	var entero1 int = 5
	var entero2 int = -5

	eq := igualdad(valor1, valor2)
	ds := distinto(valor1, valor2)

	comp1, comp2, comp3, comp4 := comparacion(entero1, entero2)
	andGate, orGate, notGate := logicGates(valor1, valor2)

	fmt.Printf("%t == %t : %t\n", valor1, valor2, eq)
	fmt.Printf("%t != %t : %t\n", valor1, valor2, ds)
	fmt.Printf("%d > %d : %t\n", entero1, entero2, comp1)
	fmt.Printf("%d < %d : %t\n", entero1, entero2, comp2)
	fmt.Printf("%d >= %d : %t\n", entero1, entero2, comp3)
	fmt.Printf("%d <= %d : %t\n", entero1, entero2, comp4)
	fmt.Printf("%t && %t : %t\n", valor1, valor2, andGate)
	fmt.Printf("%t || %t : %t\n", valor1, valor2, orGate)
	fmt.Printf("!%t : %t\n", valor1, notGate)

}
