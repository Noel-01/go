package main

import "fmt"

func main() {
	
	slice := []int{1,2,3,4}
  copia := make([]int, 4)

	//copy copia un arreglo a otro
	//IMPORTANTE: copia el minimo de elementos de ambos slices, es decir:
	// si uno tiene 3 elementos y el otro 2 solo se copiarán 2 elementos
	//Para asegurarnos de que va a funcionar bien se puede puner de la siguiente forma
	//  	slice := []int{1,2,3,4}
  //    copia := make(len(slice)) 
	copy(copia, slice)
	
	fmt.Println(slice)
	fmt.Println(copia)
	
}
