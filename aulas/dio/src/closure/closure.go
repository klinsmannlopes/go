package main

import "fmt"

// closure: funcao "chamar" uma variavel que esta em outra funcao

func main() {
	x := 0

	numero := func() int {

		x++ // incremento somar 1

		return x

	}

	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
}
