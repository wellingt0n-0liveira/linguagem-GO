//Faça um programa que receba 4 alturas usando um laço de repetição
//  e realize a soma dessas alturas.

package main

import "fmt"

func main() {

	var altura, total float64

	for i := 1; i <= 4; i++ {

		fmt.Printf("Digite a altura %d: ", i)
		fmt.Scanf("%f", &altura)

		total += altura
	}
	fmt.Printf("A soma das alturas é: %.2f\n", total)

}
