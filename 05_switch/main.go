package main

import "fmt"

func main() {

	/*"1- banana"
	  "2- maça"
	  "3- pera"
	  "4- morango"
	*/

	var opicao int

	fmt.Println("Escolha uma fruta:(1,2,3,4) ")
	fmt.Scanf("%d", &opicao)

	switch opicao {
	case 1:
		fmt.Println("Você escolheu banana")
		fmt.Println("Banana é bom para fazer vitamina")
	case 2:
		fmt.Println("Você escolheu maça")
		fmt.Println("Maçã é bom para saúde")
	case 3:
		fmt.Println("Você escolheu pera")
		fmt.Println("Pera é bom para fazer suco")
	case 4:
		fmt.Println("morango")
		fmt.Println("Morango é bom para fazer sobremesa")
	default:
		fmt.Println("Entre com uma opção inválida")
	}

}
