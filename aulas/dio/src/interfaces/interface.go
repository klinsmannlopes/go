package main

import (
	"fmt"
	"math"
)

// Interfaces sao colecoes de metodos

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

// aqui estou dizendo que a struct quadrado implementa a interface geometria - metodo area
func (q *quadrado) area() float64 {
	return q.lado * q.lado
}

// aqui estou dizendo que a struct quadrado implementa a interface geometria - metodo perimetro
func (q *quadrado) perimetro() float64 {
	return 4 * q.lado
}

// aqui estou dizendo que a struct circulo implementa a interface geometria - metodo area
func (c *circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

// aqui estou dizendo que a struct circulo implementa a interface geometria - metodo perimetro
func (c *circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// aqui estou dizendo que a struct quadrado implementa a interface geometria - metodo medir
func medir(g geometria) {
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimetro: ", g.perimetro())
}

func main() {
	q := &quadrado{
		lado: 10,
	}
	c := &circulo{
		raio: 5,
	}

	medir(q)
	medir(c)
}
