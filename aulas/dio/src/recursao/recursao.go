package main

import "fmt"

// recursao e a capacidade de uma funcao chamar ela mesma

func fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func main() {
	fmt.Println(fatorial(5))
}
