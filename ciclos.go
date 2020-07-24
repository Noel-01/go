package main

import "fmt"

// En go solo existe un bucle y es el for
func main() {
	
	// Este es un uso normal
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	
	//Se puede usar como un while
	j:=0
	for j<10 {
		fmt.Println(j)
		j++
	}
	

}
