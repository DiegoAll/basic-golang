package main

import "fmt"

func main() {

	//For condicional
	fmt.Println("For condicional")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	//For while
	fmt.Println("For while")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println()

	//For condicional (decremento)
	for i := 10; i > 0; i-- {
		fmt.Println(i)

	}

	fmt.Println()

	//For Range: Cuando se tiene una colección de un objeto.
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for i, par := range listaNumerosPares {
		fmt.Printf("posición %d número par: %d \n", i, par)

	}

	//For forever
	/*
		counterForever := 0
		for {
			fmt.Println(counterForever)
			counterForever++
		}
	*/
	/*Nota: Tener en cuenta que en Golang solo existe unciclo iterativo que es "For"-
	En otros lenguajes tienen otros ciclos como For, While, Do While, etc
	Tener en cuenta que enc asod e utilizar el ciclo For forever, debe haber una manera que dentro del
	bloque de código pueda salir del ciclo for, de lo contrario sera un gasto inoficioso de recursos y deberá
	detenerse es ciclo de forma manual. */
}
