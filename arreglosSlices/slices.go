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
func Capacidad() {
	elementos := make([]int, 5, 10)
	fmt.Println("Capacidad: ", cap(elementos))
	fmt.Println("Longitud: ", len(elementos))
	/*
	   se crea un slice de enteros con capacidad de 0 y longitud de 0.
	   El slice se inicializa con un tamaño de 0, lo que significa que no contiene elementos inicialmente.
	   Luego, se utiliza un bucle for para agregar números del 0 al 99 al slice utilizando la función append.
	   Finalmente, se imprime la capacidad y la longitud del slice.
	   El resultado muestra que la capacidad del slice es 128, ya que se han agregado 100 elementos, y la longitud es 100, ya que todos los elementos se han agregado correctamente.
	   Esto demuestra cómo se puede crear un slice dinámico en Go que puede crecer según sea necesario al agregar elementos.
	*/
	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)

	}
	fmt.Println("Capacidad: ", cap(nums))
	fmt.Println("Longitud: ", len(nums))
}
