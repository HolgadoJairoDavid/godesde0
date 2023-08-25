package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func TablaDeMultiplicar() string {
	fmt.Println("Ingrese un número: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			TablaDeMultiplicar()
		} else {
			for i := 0; i <= 10; i++ {
				texto += fmt.Sprintf("%d X %d = %d \n", numero1, i, numero1*i)
			}
			return texto
		}
	}
	return ""
}
