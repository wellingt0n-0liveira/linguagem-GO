package main
import "fmt"

func main() {
	var idade int
	fmt.Println("Entre com a idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Maior de idade")
		fmt.Println("Pode beber")
	}

		if idade <= 17 {
			fmt.Println("Menor de idade")
			fmt.Println("Nao pode beber")
		}
	}
