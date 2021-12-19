package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint
}

func main() {
	fmt.Println("Arquivo Struct")

	var u usuario
	u.nome = "Davi"
	u.idade = 4
	fmt.Println(u)

	e := endereco{"Rua x", 1}
	u2 := usuario{"Liam", 2, e}
	fmt.Println(u2)

	u3 := usuario{nome: "José"}
	fmt.Println(u3)
}
