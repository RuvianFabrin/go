package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//As chaves devem ser do mesmo tipo e os valores também
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobreNome": "Silva",
	}
	fmt.Println(usuario)
	//Para acessar
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"segundo":  "Pedro",
		},
		"facudade": {
			"primeiro": "usp",
			"segundo":  "puc",
		},
	}
	fmt.Println(usuario2)

	//deletar uma chave
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	//Adicionando
	usuario2["sig"] = map[string]string{
		"primeiro": "joao",
		"segundo":  "Pedro",
	}
	fmt.Println(usuario2)
}
