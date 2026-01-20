package main

import "fmt"

func main() {
	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	// Usando for
	// for i := 1; i < 10; i++ {
	// 	// if i%2 == 0 {
	// 	// 	fmt.Println("Par:", i)
	// 	// } else {
	// 	// 	fmt.Println("Impar:", i)
	// 	// }

	// 	// usando switch
	// 	switch i {
	// 	case 1:
	// 		fmt.Println("Um")
	// 	case 2:
	// 		fmt.Println("Dois")
	// 	case 3:
	// 		fmt.Println("Tres")
	// 	default:
	// 		fmt.Println("Outro numero")
	// 	}

	// 	switch {
	// 	case i%2 == 0:
	// 		fmt.Println("Par:", i)
	// 	default:
	// 		fmt.Println("Impar:", i)
	// 	}
	// 	//
	// }

	// loop
	// for x := 0; x < 10; x++ {
	// 	fmt.Println(x)
	// }

	// Loop Repetição Hierárquica
	// for horas := 0; horas <= 12; horas++ {
	// 	fmt.Println("Horas: ", horas)

	// 	for min := 0; min < 60; min++ {
	// 		fmt.Println("Minutos: ", min)
	// 	}
	// }

	// while
	// x := 0
	// for x < 20 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// Loop Infinito Break
	// x := 0
	// for {
	// 	if x < 20 {
	// 		x++
	// 		fmt.Println("x < 20")
	// 	} else {
	// 		break
	// 	}
	// }

	// Continue
	for x := 0; x < 20; x++ {
		if x == 3 {
			continue
		}
		fmt.Println(x)
	}

}
