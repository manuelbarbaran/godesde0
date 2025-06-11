package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Iniciando la funcion VemosDefer")
	defer fmt.Println("Finalizando la funcion VemosDefer")
	fmt.Println("Este es el segundo mensaje de la funcion VemosDefer")
}
func EjemploPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatal("Recuperado de un panic:", reco)
		}
	}()

	fmt.Println("Iniciando la funcion EjemploPanic")
	defer fmt.Println("Finalizando la funcion EjemploPanic")
	panic("Esto es un panic")
	//fmt.Println("Este mensaje no se imprimira debido al panic")
}
