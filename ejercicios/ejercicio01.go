package ejercicios

import "strconv"

func PrimerEjercicio(str string) (int, string) {
	number, _ := strconv.Atoi(str)

	if number > 100 {
		return number, "Es mayor a 100"
	} else {
		return number, "Es menor a 100"
	}
}
