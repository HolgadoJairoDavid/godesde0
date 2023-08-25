package main

import (
	"fmt"

	"github.com/HolgadoJairoDavid/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(15)
	fmt.Println(estado, texto)
}
