package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al escribir archivo " + err.Error())
		return
	}

	fmt.Fprint(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaDeMultiplicar()

	if !Append(texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(texto string) bool {
	arch, err := os.OpenFile(filename, os.O_RDONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error abriendo el archivo " + err.Error())
		return false
	}

	_, erro := arch.WriteString(texto)

	if erro != nil {
		fmt.Println("Error escribiendo archivo " + erro.Error())
		arch.Close()
		return false
	}
	arch.Close()
	return true
}

func LeerArchivo() {
	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error abriendo el archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))

}

func LeerArchivo2() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error abriendo el archivo")
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()

}
