package funciones

import "fmt"

func Calculos() {
	var numero3 int = 2
	var numero4 int = 4

	suma := func(n1, n2 int) int {
		return n1 + n2 + numero3 + numero4
	}

	fmt.Println(suma(10, 25))

	suma = func(n1, n2 int) int {
		return n1 * n2 * numero3 * numero4
	}

	fmt.Println(suma(10, 25))
}
