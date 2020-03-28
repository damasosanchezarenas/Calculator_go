package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calcu struct {
}

func (calcu) operate(entrada string, operador string) int {
	//Hacemos un split, para sacar ambos digitos. Devuelve un array de dos operadores (strings).
	entradaLimpia := strings.Split(entrada, operador)

	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	default:
		fmt.Println(operador, "No esta soportado, se te devolverá un -1")
		return -1
	}
}

func parsear(entrada string) int {
	//Atoi es un casteo para pasar string--int devuelve dos valores. Si crearamos otra variable en vez de _, tendriamos que usarla si o si
	operador1, _ := strconv.Atoi(entrada)
	return operador1
}

func leerEntrada() string {

	scanner := bufio.NewScanner(os.Stdin) //Declaración del scanner
	scanner.Scan()
	operacion := scanner.Text() //Coge lo que nosotros le digamos por teclado.

	return operacion
}

func main() {

	entrada := leerEntrada()
	operador := leerEntrada()

	calculadora := calcu{}
	fmt.Println(calculadora.operate(entrada, operador))
}
