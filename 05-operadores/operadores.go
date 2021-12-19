package main

import "fmt"

func main() {
	//Aritimeticos
	soma := 1 + 2
	sub := 1 - 2
	div := 1 / 2
	mult := 1 * 2
	resto := 1 % 2

	fmt.Println(soma, sub, div, mult, resto)

	//Não é premitido fazer operações com tipo diferentes
	//int32 e int16 por exemplo
	//Fim dos aritiméticos

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	//1>2
	//>=
	// ==
	// <=
	// !=
	// FIM DOS OPERADORES RELACIONAIS

	//OPERADORES LÓGICOS
	// && = E
	// || = OU
	// ! = INVERTE (TRUE PARA FALSE, FALSE PARA TRUE)

	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNARIOS
	// numeros:=10
	// numero++ //Soma 1
	// numero+=15 //soma 15 ao valor anterior
	// não existe -- ou ++ antes (++numero)
	// o mesmo para o (menos) -
	// para outros
	// numero *=3 //multiplica o valor anterior por 3
	// numero /=3 //divide o valor anterior por 3
	// numero %=3 //resto da divisão do valor anterior por 3

	//FIM OPERADORES UNARIOS

	// Não tem operador ternario Ex.: x=(true)?true:false;
}
