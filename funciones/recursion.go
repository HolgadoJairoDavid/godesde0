package funciones

import "fmt"

func Exponencia(n int) {
	if n > 100 {
		return
	}
	fmt.Println(n)
	Exponencia(n * 2)
}
