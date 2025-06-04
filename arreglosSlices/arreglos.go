package arreglosslices

import "fmt"

var tabla [10]int = [10]int{0, 52, 63, 56}
var matriz [10][15]int

func MuerstroArreglos() {

	tabla[7] = 33
	tabla[9] = 20
	tabla2 := [10]string{"Manuel", " ", "Barbaran"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[8][5] = 25
	fmt.Println(matriz)

}
