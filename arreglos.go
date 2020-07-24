package main

import "fmt"

func main() {
	
	var arreglo [10] int //10 espacios para enteros, por defecto 0s
	var arreglo2 [10] string //10 espacios para palabras o frases
	
	fmt.Println(arreglo)
	fmt.Println(arreglo2)
	
	
	arreglo3 := [3] int {1,2,3}
	arreglo4 := [4] int {1}
	
	fmt.Println(arreglo3)
	fmt.Println(len(arreglo4)) //tamaño del arreglo
	
	for i := 0; i < len(arreglo3); i++ {
		fmt.Println(arreglo3[i])
	}

}
