package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	//Inferencia de tipo
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	println(variavel5, variavel6)

	//Inverter valores de variavel
	variavel5, variavel6 = variavel6, variavel5

	println(variavel5, variavel6)
}
