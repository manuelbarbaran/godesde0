package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumero() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese Numero 1 : ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese Numero 2 : ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese Leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

}

func IngresoNumeroMultiplicar() int {

	numero := 0
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese numero de la tabla de multiplicar")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			break
		}
	}
	return numero
}
