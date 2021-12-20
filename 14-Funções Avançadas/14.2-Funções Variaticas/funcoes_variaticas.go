package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//Só pode ter um valor variatico, e tem que ser por ultimo
func escrever(texto string, numeros ...int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println(texto, ": ", total)
}
func main() {
	fmt.Println("Funções variaticas")
	totalDaSoma := soma(1, 3, 4, 5, 6, 7, 8, 9, 20, 3, 3, 2, 1)
	fmt.Println(totalDaSoma)

	escrever("Total", 3, 4, 5, 6, 7, 8, 9, 20, 3, 3, 2, 1)
}
