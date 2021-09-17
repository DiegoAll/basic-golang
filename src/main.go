package main

import (
	"bufio"
	"fmt"
	"os"
)

func oddOrEvent(num int) string {
	if num%2 == 0 {
		return "El número es par"
	}
	return "El número es impar"
}

func validatepassword(pw string) {
	pwTrue := "JoseNest0rPekerman"

	if pw != pwTrue {
		fmt.Println("Acceso denegado. Contraseña inválida")
	} else {
		fmt.Println("Acceso autorizado. Bienvenido")
	}
}

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With AND
	if valor1 == 1 && valor2 == 3 {
		fmt.Println(("Es verdad"))
	}

	//With OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	numeroPar := 8
	numeroImpar := 33

	fmt.Println(oddOrEvent(numeroPar))
	fmt.Println(oddOrEvent(numeroImpar))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	password := scanner.Text()

	validatepassword(password)
}
