package defer_panic

import (
	"fmt"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es un segundo mensaje")
}
