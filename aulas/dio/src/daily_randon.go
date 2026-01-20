package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	devs := []string{
		"lin",
		"bury",
		"guedes",
		"vinicius",
		"henrique",
		"randerson",
		"klinsmann",
	}

	rand.Shuffle(len(devs), func(i, j int) {
		devs[i], devs[j] = devs[j], devs[i]
	})

	fmt.Println("Random dev order for the day:")
	for _, dev := range devs {
		fmt.Printf("%s\n", dev)
	}
}
