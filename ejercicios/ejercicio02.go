package ejercicios

import (
	"fmt"

	"github.com/godesde0/teclado"
)

func TablaDeMultiplicar() string {

	var tablaCompleta string
	numero := teclado.IngresoNumeroMultiplicar()

	for i := 1; i <= 10; i++ {
		tablaCompleta += fmt.Sprintf("%d X %d = %d\n", numero, i, (numero * i))
	}
	return tablaCompleta
}
