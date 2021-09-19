package main

import "fmt"

func main() {

	//Array [Son valores inmutables] No se puede agregar otro elemento pero si se pueden editar los valores
	var array [4]int
	fmt.Println(array)

	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Println("La longitud del array es:", len(array))
	fmt.Println("La capacidad del array es:", cap(array))

	//Slices Similares a los arrays, pero no se le indica el tamaño que va a tener.
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println("La longitud del slice es:", len(slice))
	fmt.Println("La capacidad del slice es:", cap(slice))

	//Slicing es lo que se utiliza al momento de manejar arrays,slices o listas para poder interactuar con cada uno de sus elementos.
	//Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  //El indice que va despupes de los 2 puntos es esclusivo
	fmt.Println(slice[2:4]) //Lim inferior es inclusivo pero lim superior es esclusivo
	fmt.Println(slice[4:])  //Primer indice es inclusivo

	//Adicionar elementos en los slices
	slice = append(slice, 7)
	fmt.Println(slice)

	//Adicionar una lista en el slice
	newSlice := []int{8, 9, 20}
	slice = append(slice, newSlice...)
	/*La extension agrega los ..., indica que se realiza la descompresión y en vez de agregar una nueva lista,
	agrega cada uno de los elementos de forma independiente */
	fmt.Println(slice)

}
