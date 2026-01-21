// Faça um programa que receba um número inteiro
// e calcule sua raiz quadrada e exiba o resultado.
package main

import "math"
import "fmt"

func main() {

	var numero int
	fmt.Println("Digite um número :")
	fmt.Scanf("%d", &numero)

	resultado := math.Sqrt(float64(numero))
	fmt.Printf("A raiz quadrada de %d é %.2f\n", numero, resultado)
}
