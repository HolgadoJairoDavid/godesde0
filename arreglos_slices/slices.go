package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 3, 4, 5}
var arreglo [10]int = [10]int{6, 7, 8, 9, 45, 6, 7, 8, 8, 9}

func MuestroSlices() {
	fmt.Println(tablaSlice)
	porcion := arreglo[3:]
	porcion2 := arreglo[:4]
	porcion3 := arreglo[4:5]
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, capacidad %d", len(nums), cap(nums))
}
