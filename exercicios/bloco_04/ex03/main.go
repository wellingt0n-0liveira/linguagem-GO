//Faça um programa que receba uma quantidade indefinida de valores
//  correspondentes a “saldo em conta”, mas quando o usuário apertar
//  “enter” sem digitar valor algum, o programa para de receber valores,
//  e exibe a soma de todos os valores digitados anteriormente.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var soma float64

	for {
		var entrada string
		fmt.Print("Digite o saldo em conta (ou aperte 'enter' para sair): ")

		// Lê linha inteira, ignorando espaços
		fmt.Scanf("%s\n", &entrada) // ← Adiciona \n para consumir o Enter

		if entrada == "" {
			break
		}

		valor, erro := strconv.ParseFloat(entrada, 64)
		if erro != nil {
			fmt.Println("Valor inválido. Por favor, digite um número válido.")
			continue
		}

		soma += valor
	}

	fmt.Printf("A soma dos saldos é: %.2f\n", soma)
}
