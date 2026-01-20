package main
import "fmt"

func main() {
	var idade int
	fmt.Println("Entre com a idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 66 {

		fmt.Println("Boa idade para descansar")
		fmt.Println("manere na bebedeira")

	} else if idade >= 18 {

			fmt.Println("Maior de idade")
			fmt.Println("Pode beber")

	} else {	

		fmt.Println("Menor de idade")
		fmt.Println("Nao pode beber")

		
		}
	}
