package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("teste Pacotes")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("devbook_gmail.com")
	fmt.Println(erro)

}
