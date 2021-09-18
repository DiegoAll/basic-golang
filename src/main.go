package main

import "fmt"

func main() {

	//Defer va a ejecutar la última función antes de que todo termine, no se limita a un comando tambien puede especifricarse uan función
	// En que casos usas defer? Ej. Alr ealizar una conexión a una BD, al final para que realice el cierre de la conexión.
	//En teoria se puedne utilizar varios defer en una fucnión, ṕero la buena práctica indica que sedebe utilizar solo un defer por función.

	defer fmt.Println("Hablalo mostro")
	fmt.Println("Buenas tardes estimado")

	//continue y break
	//continue se utilzia cuando se cumple una condición dada  dentro del ciclo for puede ser algo de intereres que continue en ejecución.
	/*Por ejemplo que suceda un error controlado en el cual aunque suceda ese error, a pesar de que sucveda este error se parsea se ejecuta
	otro bloque de código pero quiero que continue la ejecución del ciclo For*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//Break
		/* Se utiliza cuando sucede algo determinado una condición determinada y se espera que a pesar de esa condición no se desea que
		el código se ejecute*/

		if i == 8 {
			fmt.Println("Break")
			break
		}

	}

}
