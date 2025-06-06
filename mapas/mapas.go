package mapas

import "fmt"

func MostrarMapas() {
	// Definición de un mapa vacío
	var mapaVacio map[string]int
	fmt.Println("Mapa vacío:", mapaVacio)

	paises := make(map[string]string)
	paises["ES"] = "España"
	paises["AR"] = "Argentina"
	paises["CL"] = "Chile"
	paises["MX"] = "México"
	paises["CO"] = "Colombia"

	fmt.Println("Mapa de países:", paises)
	// Iteración sobre un mapa
	for clave, valor := range paises {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
	// Verificación de la existencia de una clave
	if valor, existe := paises["ES"]; existe {
		fmt.Println("El valor para la clave 'ES' es:", valor)
	} else {
		fmt.Println("La clave 'ES' no existe en el mapa.")
	}
	// Eliminación de una clave del mapa
	delete(paises, "AR")
	fmt.Println("Mapa de países después de eliminar 'AR':", paises)

	// Inicialización de un mapa con valores
	mapa := map[string]int{
		"uno":    1,
		"dos":    2,
		"tres":   3,
		"cuatro": 4,
	}

	// Acceso a un valor en el mapa
	valor := mapa["dos"]

	// Verificación de la existencia de una clave
	if valor, existe := mapa["cinco"]; existe {
		fmt.Println("El valor es:", valor)
	} else {
		fmt.Println("La clave 'cinco' no existe en el mapa.")
	}

	fmt.Println("Valor de 'dos':", valor)
}
