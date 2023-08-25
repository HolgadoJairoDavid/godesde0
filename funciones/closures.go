package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarCloure() {
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}
}
