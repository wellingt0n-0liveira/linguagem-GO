// Faça um programa que conte quantas vezes a letra “a” aparece em uma palavra
package main

import (
	"fmt"
	"strings"
)

func main() {

	var palavra string
	fmt.Print("Digite uma palavra: ")
	fmt.Scanf("%s", &palavra)

	palavraLower := strings.ToLower(palavra)

	contador := 0

	for _, letraByte := range palavraLower {

		if string(letraByte) == "a" {
			contador++
		}
	}

	fmt.Printf("A letra 'a' aparece %d vezes na palavra '%s'.\n", contador, palavra)
}
