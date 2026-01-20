package main

// panic: e uma forma de parar a execucao do programa
// Recover: e uma forma de recupar a execucao do programa

import "fmt"

func main() {
	defer func() {
		x := recover()
		fmt.Println(x)
		fmt.Println("Continuando a execução...")
	}()
	panic("PANIC")

}
