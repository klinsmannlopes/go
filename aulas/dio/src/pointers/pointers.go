package main

import "fmt"

func main() {
	// 1. Variável normal
	x := 10
	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x na memória:", &x) // & pega o endereço

	// 2. Criando um ponteiro
	// Um ponteiro é uma variável que guarda um endereço de memória
	var ponteiro *int = &x

	fmt.Println("\nValor do ponteiro (endereço de x):", ponteiro)
	fmt.Println("Valor armazenado nesse endereço (*ponteiro):", *ponteiro) // * lê o valor (Dereferencing)

	// 3. Modificando valor através do ponteiro
	*ponteiro = 20 // Vai lá no endereço e muda o valor
	fmt.Println("\nAlterei via ponteiro. Novo valor de x:", x)

	// 4. Passagem de valor vs referência em funções
	valor := 5
	mudarValor(valor) // Passa CÓPIA
	fmt.Println("\nTentativa de mudar valor (cópia):", valor)

	mudarReferencia(&valor) // Passa ENDEREÇO
	fmt.Println("Tentativa de mudar referência (ponteiro):", valor)
}

func mudarValor(n int) {
	n = 100
}

func mudarReferencia(n *int) {
	*n = 100
}
