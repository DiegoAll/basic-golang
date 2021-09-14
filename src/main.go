package main

import "fmt"

func main() {

	//Funciones mas importantes del paquete fmt

	//Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	expresion := "Parcero"
	cantidad := 500
	fmt.Printf("%s por favor me da %d de cilantro\n", expresion, cantidad)

	//Cuando no se sabe el tipo de valor, la buena práctica es intentar asignar el tipo de dato al que pertenece
	fmt.Printf("%v por favor me da %v de cilantro\n", expresion, cantidad)

	//Sprintf va a generar un string pero no lo va a imprimir en consola pero lo va a guardar como un string
	message := fmt.Sprintf("%v por favor me da %v de cilantro", expresion, cantidad)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)

	fmt.Printf("Cantidad: %T\n", cantidad)

	//Las anteriores son las funciones mas "importantes" o útiles del paquete fmt.

}
