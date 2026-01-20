package main

import "fmt"

func main() {

	// colecao ordenada (lista) de pares chave-valor
	// var x map[string]int
	// x e um mapa de  strings para inteiros

	// x := make(map[string]int)
	// x["chave"] = 10

	// x := make(map[int]int)
	// x[1] = 20
	// x[2] = 30

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"

	fmt.Println(elemento["H"])
	fmt.Println(elemento)
}
