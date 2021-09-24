package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	//Si no tuvieramops construía la interfaz para calcular el area sería de la siguiente forma:  myCuadrado.area() y myRectangulo.area()

	// Cuando se tienen varias estructuras de datos que comparten la misma función al implementar interfaces es mucho mas eficiente
	//Pero como disponemos de la interface la utilizaremosd e la siguiente forma:

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista de interfaces
	//En otros lenguajes se pueden tener listas con  elementos de diversos tipos de datos, en GO no se puede hacer
	// Cuando se crea el tipo de dato sea slice o array se debe indicar el tipo de dato que se va a insertar allí
	// hay una forma de simularlo y es:

	//Un slice de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
