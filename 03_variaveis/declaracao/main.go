package main

import "fmt"

func main() {
	var x int
	fmt.Println("Valor de x =", x)

	var y int = 10
	fmt.Println("Valor de y =", y)

	x = 45
	fmt.Println("Novo valor de x =", x)

	var z, w, v float64
	fmt.Printf("z = %f | w = %f | v = %f \n", z, w, v)

	var a, b, c string = "Go", "é", "legal"
	fmt.Println(a, b, c)

}
