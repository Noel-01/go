package main

import "fmt"

func main(){

	var x,y *int  // Punteros
	entero := 5   // Entero normal 
	
	x = &entero   // Puntero apuntando a la dirección de memoria del entero
	y = &entero   // Puntero apuntando a la dirección de memoria del entero
	
	*x = 6 // Cambio el valor de lo que hay en la dirección de memoria a la que apunta el puntero *x
	
	/* Imprime la dirección de memoria */
	fmt.Println(x)
	fmt.Println(y)
	
	/* Imprime el valor de lo que hay en la dirección de memoria */
	fmt.Println(*x)
	fmt.Println(*y)

}
