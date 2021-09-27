package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {

	c := make(chan string, 2)

	c <- "Message 1"
	c <- "Message 2"

	fmt.Println(len(c), cap(c))

	//Range y close
	/*Decirle al runtime de GO que va a cerrar el canal, eso quiere decir que ese canal no va a recibir otro dato adicional a pesar de que tenga
	mas capacidad. Lo ideal es cerrar los canales una vez ya dejan de usarse.
	*/

	//Range y close
	close(c)

	//c <- "Message 3"

	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("message1", email1)
	go message("message2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
