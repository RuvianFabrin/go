package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	/*i := 0
	for i < 3 {
		i++
		fmt.Println("Incrementando i.")
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	for j := 0; j < 2; j++ {
		fmt.Println("Incrementando J= ", j)
		time.Sleep(time.Second)
	}*/

	//Com slice é a mesma coisa
	/*nomes := [3]string{"Jose", "Lucas", "Marco"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	//Sem usar o indice
	for _, nome := range nomes {
		fmt.Println(nome)
	}*/
	//Iterar string
	/*for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}*/

	//Iterar Map
	usuario := map[string]string{
		"teste":  "ok",
		"teste2": "ok2",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Não dá pra fazer range em struct

	//Loop infinito
	for {
		fmt.Println("Fica Infinitamente")
		time.Sleep(time.Second)
	}
}
