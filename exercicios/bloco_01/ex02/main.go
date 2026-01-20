// faça um programa que exiba as tabelas verdade com expressões && e ||

package main

import "fmt"

func main() {
	fmt.Println("Tabelas verdade de &&")
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && true =", false && true)
	fmt.Println("false && false =", false && false)

	fmt.Println("Tabelas verdade de ||")
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || true =", false || true)
	fmt.Println("false || false =", false || false)

}
