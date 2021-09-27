package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	/*
		Los channels te permiten compartir datos entre las goroutines ya que ello maneja de forma nativa la comunicación entre ellos
		además de otros datos primitivos como por ejemplo los WaitGroups para manejar la concurrencia
	*/

	/*El segundo parámetro corresponde a la cantidad limite de goroutines que recibirá el channel, es buena practica especidifcarlo, en caso de que no se
	especifique el va a tomar un valor dinamico en todo momento.
	Así como esta solo va a recibir una goroutine de tipo string.
	*/

	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)

}
