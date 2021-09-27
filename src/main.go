package main

import (
	"fmt"
	"sync"
	"time"
)

/*func say(text string) {
	fmt.Println(text)
}*/

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {

	/*Este paquete "sync" permite interactuar de forma primitiva con las Goroutines, lo cual lo hace mas eficiente
	sin embargo implica un costo mas alto para ser mantenido
	*/
	var wg sync.WaitGroup

	fmt.Println("Que mas pues")
	wg.Add(1)

	go say("panita", &wg) //para empezar a manejar las goutines solo se utiliza el keyword "go"

	wg.Wait()

	/*Las Go routines se utilizan mucho con funciones anónimas
	¿Qué es una funcion anonima?
	Es una función que no es declarada anteriormente, sino que se realiza dentro de la misma función y tiene su vida en ella misma.
	*/
	go func(text string) {
		fmt.Println(text)
	}("Bye parcero")

	time.Sleep(time.Second * 1)
}
