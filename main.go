package main

import (
	"github.com/godesde0/ejer_interfaces"
	"github.com/godesde0/modelos"
)

func main() {

	/*
					estado, texto := variables.ConviertoaTexto(123456)
					fmt.Println(estado)
					fmt.Println(texto)

					if os := runtime.GOOS; os == "linux" || os == "OS X." {

						fmt.Println("Esto no es Windows, es ", os)
					} else {
						fmt.Println("Esto es windows")

					}

					switch os := runtime.GOOS; os {
					case "linux":
						fmt.Println("Esto es Linux")
					case "darwin":
						fmt.Println("Esto es Darwin")
					default:
						fmt.Printf("%s \n", os)
					}

				valorEntero, mensaje := ejercicios.VerificarEntero100("102")

				fmt.Println("El valor entero es ", valorEntero)
				fmt.Println(mensaje)

			teclado.IngresoNumero()
			iteraciones.Iterar()


		fmt.Println(ejercicios.TablaDeMultiplicar())
	*/
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeerArchivo2()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)
	//arreglosslices.Capacidad()
	//mapas.MostrarMapas()
	//usuarios.AltaUsuario()
	Manuel := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirando(Manuel)

	Nicol := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Nicol)

}
