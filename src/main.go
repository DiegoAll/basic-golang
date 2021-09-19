package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i]) // Se realiza una conversion a string ya que retorna un tipo de dato byte correspondiente al codigo ascii del caracter
	}

	if text == textReverse {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}
}

/*Ejemplos de palindromos
-Anita lava la tina
-Amor a Roma
-Allí ves a Sevilla
-Amigo no gima
-Amad a la dama
-Ella te da detalle
-Anula la luna
-Isaac no ronca asi
-Ligar es ser agil
-Luz azul
-Oso baboso
-Se van sus naves
-Sometamos o matemos
-Ojo rojo
-Son mulas o civicos alumnos
-Son robos o sobornos
-Amor a roma
*/

func main() {

	slice := []string{"Que", "mas", "pues"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	fmt.Println()

	//Sino me interesa el indice, solamente el valor, lo escapamos.
	for _, valor := range slice {
		fmt.Println(valor)
	}
	palabra := "amigonogima"
	isPalindromo(palabra)

	//Case insensitive
	//Ama
	var palabra2 string
	fmt.Scan(&palabra2)
	word := strings.ToLower(palabra2)

	isPalindromo(word)

}
