package main

import "fmt"

func main() {
	fmt.Println("O que vc quiser!!")

	fmt.Println(`Aqui a gente
	pode escrever
	o que quiser
	sem problemas
	em multiplas linhs de código usando a crase`)

	fmt.Println("Teodoro tem ", len("Teodoro"), " letras.")

	fmt.Println("Teodoro"[0]) // retorna o valor em byte (uint8) da tabela ASCII

	fmt.Println(string("Teodoro"[0])) // converte o valor em byte para string
	fmt.Println(string("Teodoro"[1]))
	fmt.Println(string("Teodoro"[2]))
	fmt.Println(string("Teodoro"[3]))
	fmt.Println(string("Teodoro"[4]))
	fmt.Println(string("Teodoro"[5]))

}
