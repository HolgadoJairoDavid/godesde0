package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises)
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  56,
		"Chivas":       37,
		"Boca Juniors": 89,
	}
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipos %s, tiene un puntaje de %d \n", equipo, puntaje)
	}
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Println(puntaje, existe)
}
