package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //Declaración del scanner
	scanner.Scan()

	operacion := scanner.Text() //Coge lo que nosotros le digamos por teclado.

	operador := "p"
	valores := strings.Split(operacion, operador) //Hacemos un split, para sacar ambos digitos. Devuelve un array de dos operadores (strings).

	//Atoi es un casteo para pasar string--int devuelve dos valores. Si crearamos otra variable en vez de _, tendriamos que usarla si o si
	operador1, err1 := strconv.Atoi(valores[0])
	operador2, err2 := strconv.Atoi(valores[1])

	if err1 == nil && err2 == nil {

		switch operador {
		case "+":
			fmt.Println(operador1 + operador2)
		case "-":
			fmt.Println(operador1 - operador2)
		case "*":
			fmt.Println(operador1 * operador2)
		case "/":
			fmt.Println(operador1 / operador2)
		default:
			fmt.Println(operador, "No esta soportado")
		}
	} else {
		fmt.Println("Asegurese de introducir correctamente la operación.")
	}

}
