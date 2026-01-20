package main

import "fmt"

// array
func main() {
	// var x [10]int
	// x[4] = 80
	// fmt.Println(x)

	// var x [5]float64
	// x[0] = 5.3
	// x[1] = 8
	// x[2] = 4.2
	// x[3] = 2.1
	// x[4] = 7.8

	// var total float64 = 0
	// //var tamanho int32 = len(x)
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }

	// fmt.Println(total / float64(len(x)))

	// pegar fatia de um array
	// var x []float64
	// fatia := make([]float64, 4)
	// fatia := [low:high]
	// fatia := arr[0:5]

	// aqui estou printando o array, um pedaco dele, colocando o inicio e o fim
	// arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	// fatia := arr[0:3]

	// aqui estou adicionando mais valores ao array
	// fatia1 := []int{1, 2, 3}
	// fatia2 := append(fatia1, 4, 5, 6)

	fatia1 := []int{1, 2, 3}
	fatia2 := make([]int, 2)
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}
