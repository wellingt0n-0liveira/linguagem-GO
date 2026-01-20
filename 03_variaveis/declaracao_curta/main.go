package main

import "fmt"

func main() {

	x := 45 // isso é um inyteiro
	nome := "Golang"
	y := 0.45

	println(x, nome, y)

	x = 90
	fmt.Println("Novo valor de x =", x)

	fmt.Printf("Tipo de x = %T \n", x)
	fmt.Printf("Tipo de nome = %T \n", nome)
	fmt.Printf("Tipo de y = %T \n", y)

}
