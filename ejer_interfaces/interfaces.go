package ejer_interfaces

import (
	"fmt"

	"github.com/godesde0/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {

	humano.Respirar()

	fmt.Printf("soy un humano/a de sexo %s\n", humano.Sexo())
	fmt.Println("Estoy respirando")
}
