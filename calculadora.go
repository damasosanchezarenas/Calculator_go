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

	operacion := scanner.Text()              //Coge lo que nosotros le digamos por teclado.
	valores := strings.Split(operacion, "+") //Hacemos un split, para sacar ambos digitos. Devuelve un array de dos operadores (strings).

	operador1, _ := strconv.Atoi(valores[0]) //Atoi devuelve dos valores.
	operador2, _ := strconv.Atoi(valores[1])

	fmt.Println(operador1 + operador2)

}
