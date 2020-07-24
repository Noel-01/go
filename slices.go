package main

import "fmt"

// Los slices son como los arreglos sin definir el tamaño al principio

func main() {

	matriz := [] int {1,2,3,4}
	
	fmt.Println(matriz)
	
	matriz2 := [] int {1,2}
	
	if matriz2 == nil {
		fmt.Println("Está vacio")
	}else{
	  fmt.Println(len(matriz2))
	}
	
	arreglo := [3] int {1,2,3}
	slice := arreglo[0:2] // Eso lo que hace es cocer del 0 al 2 (inclusive)
	fmt.Println(slice) //Imprimer [1 2]
	
	
}
