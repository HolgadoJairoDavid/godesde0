package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func TablaDeMultiplicar() {
	fmt.Println("Ingrese un número: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			TablaDeMultiplicar()
		} else {
			for i := 0; i <= 10; i++ {
				fmt.Println(numero1 * i)
			}
		}
	}
}
