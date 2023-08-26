package ejer_interfaces

import (
	"fmt"

	"github.com/HolgadoJairoDavid/godesde0/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirndo \n", hu.Sexo())
}
