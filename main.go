package main

import (
	"fmt"

	"github.com/godesde0/ejercicios"
)

func main() {

	valorEntero, mensaje := ejercicios.VerificarEntero100("102")

	fmt.Println("El valor entero es ", valorEntero)
	fmt.Println(mensaje)

}
