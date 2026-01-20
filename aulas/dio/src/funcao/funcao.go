package main

import "fmt"

// funcao tbm e chamada de procedimento e sub-rotina
// ela pega um dado de entrada e da um dado de saida

func media(lista []float64) float64 {
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista)) // return interrompe a funcao
}

func main() {
	// programa que calcula a media de uma sala
	lista := []float64{98, 93, 77, 82, 83}
	media := media(lista)
	fmt.Println(media)
}
