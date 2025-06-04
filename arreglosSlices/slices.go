package arreglosslices

import "fmt"

var tablaS []int = []int{2, 5, 6}
var arreglo [10]int = [10]int{8, 9, 6, 2, 7, 8, 4, 55}

func MostrarSlices() {

	fmt.Println(tablaS)

	porcion := arreglo[3:]
	porcion2 := arreglo[:5]
	porcion3 := arreglo[5:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}
