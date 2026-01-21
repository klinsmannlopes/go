package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Cria o arquivo "texto.txt". Se já existir, ele recria zerado.
	f, err := os.Create("texto.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // Importante: fecha o arquivo quando a main terminar

	// Escreve no arquivo 'f' em vez de 'os.Stdout'
	if _, err := io.WriteString(f, "Bom dia"); err != nil {
		log.Fatal(err)
	}
}
