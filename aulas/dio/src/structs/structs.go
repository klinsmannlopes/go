package main

import "fmt"

// sao colecoes de campos
type pessoa struct {
	nome  string
	idade int
}

// metodo e uma funcao associada a um tipo particular
// Isto e, em POO, objeto e um valor (varieavel) e metodo e uma funcao associada a esse objeto
type retangulo struct {
	comprimento, altura int
}

// este metodo area possui um tipo retangulo
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2 * (r.comprimento + r.altura)
}

func main() {

	// p := pessoa{
	// 	nome:  "Klinsmann",
	// 	idade: 25,
	// }

	// fmt.Println(p)
	// fmt.Println(p.nome)
	// fmt.Println(p.idade)
	r := retangulo{
		comprimento: 10,
		altura:      5,
	}

	fmt.Println("Area: ", r.area())
	fmt.Println("Perimetro: ", r.perimetro())

}
