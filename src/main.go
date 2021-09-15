package main

import "fmt"

func normalFunction() {
	fmt.Println("Hola mundo")
}

func normalFunction2(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2

}

func main() {

	//Invocar la funciión

	/*fmt.Println("Hola mundo")
	fmt.Println("Hola mundo 2")
	fmt.Println("Hola mundo 3")*/

	normalFunction()
	normalFunction2("Que mas pues parcero")
	tripleArgument(1, 3, "guineos")

	value := returnValue(3)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)

	value3, _ := doubleReturn(3)
	fmt.Println("value1", value3)

}
