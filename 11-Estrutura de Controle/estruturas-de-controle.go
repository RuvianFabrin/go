package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//if init, limita ao escopo do if
	if outronumero := numero; outronumero > 0 {
		fmt.Println("maior que zero")
	} else {
		fmt.Println("menor que zero")
	}
}
