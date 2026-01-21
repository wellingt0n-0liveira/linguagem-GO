//Faça um programa que exiba o dobro de um número inserido pelo usuário.
package main

import "fmt"

func main() {

	var x float64

	fmt.Print("Digite um número: ")
	fmt.Scanf("%f", &x)

	res := x * 2
	fmt.Printf("O dobro de %.2f é %.2f\n", x, res)
}
