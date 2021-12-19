package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array2)

	array3 := [...]string{"aa", "bb", "cc", "dd", "ee"}
	fmt.Println(array3)

	//Slice não tem tamanho definido
	slice := []int{1, 2, 3, 5, 6, 8}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	//Colocando valores no slice
	slice = append(slice, 30)
	fmt.Println(slice)

	//Pegar um pedaço
	slice2 := array2[1:3]
	fmt.Println(slice2)

	//funciona como um ponteiro
	array2[1] = "Posição alterada"
	fmt.Println(slice2)
}
