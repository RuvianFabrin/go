package main

import (
	"errors"
	"fmt"
)

func main() {
	//-==========================================
	//Tipos de int (235)
	//int8, int16, int32, int64
	//se você não declarar ( int ), ele vai usar a arquetetura do computador
	//Processador 32 int32 | processador 64 int64
	//uint = int sem sinal (Aumenta o dobro de tamanho, sem negativo)
	//alias
	//rune = int32
	//byte = int8
	//-==========================================

	//-==========================================
	//float32, float64 (22.36)
	//-==========================================

	//-==========================================
	//string ("Texto qualquer")
	//-==========================================

	//-==========================================
	//char ('T') | Mostra um numero da tabela ASCII
	//-==========================================

	//-==========================================
	//para o go todos os tipos numero começa com 0 se não indicado o valor
	//String começa com limpo
	//bool começa com false
	//error começa com nil
	//-==========================================

	//-==========================================
	//bool (true,false)
	//-==========================================

	//-==========================================
	//error (true,false)
	//-==========================================

	//Como fazer um erro no go
	var erro error = errors.New("Criou uma mensagem de erro interno")
	fmt.Println(erro)
}
