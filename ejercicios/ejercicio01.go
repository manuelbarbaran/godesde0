package ejercicios

import "strconv"

func VerificarEntero100(parametro string) (int, string) {

	var mensaje string
	parametroEntero, _ := strconv.Atoi(parametro)

	if parametroEntero > 100 {

		mensaje = "Es mayor a 100"
	} else {
		mensaje = "Es menor a 100"
	}

	return parametroEntero, mensaje
}
