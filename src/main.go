package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Faustino"] = 45
	m["Justino"] = 24

	fmt.Println(m)

	//recorrer el map
	/* Tener en cuenta que al realizar el recorrido con map se hace de forma aleatoria, no va a retornar la clave y valor en el mismo
	orden en el cual fue ingresad, como sucede de forma concurrente va a ejecutarse de forma aleatoria, Si se requiere la información en el mismo orden se
	recomiendan los slices */
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Como obtener uno de los valores
	value, ok := m["Justino"]
	fmt.Println(value, ok)
}
