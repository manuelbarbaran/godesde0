package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {

	estado, texto := variables.ConviertoaTexto(123456)
	fmt.Println(estado)
	fmt.Println(texto)
}
