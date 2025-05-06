package ejercicios

import (
	"fmt"

	"github.com/godesde0/teclado"
)

func TablaDeMultiplicar() {

	numero := teclado.IngresoNumeroMultiplicar()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", numero, i, (numero * i))
	}
}
