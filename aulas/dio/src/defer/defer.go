package main

import "fmt"

// defer: escalona nossas funcoes para serem executadas mais tarde

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda")
}

func main() {
	defer dia1()
	dia2()
}
