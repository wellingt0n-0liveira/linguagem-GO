package main

import "fmt"

func main() {

	// rodar para sempre
	for i := 1; i <= 1000; i++ {
		if 1%25 == 0 {
			fmt.Println("dane-se", i)
			fmt.Println("multiplo de 25 achado")
			break
		}
	}
}
