package main

import "fmt"

func main() {
	
	//Make crea un slice con longitud 3 y valor 0 y una capacidad de 5
	slice := make([]int, 3,5)
	
	slice[2] = 1 //Con esto puedes cambiar uno de los 3 valores de la longitud
	slice = append(slice, 2) //Se añade un numero mas, ahora es longitud 4

	//Aunque se sobre pase la longitud inicial, append reconstruye un nuevo slice
	// Si la capacidad es mas grande que la longitud y se utiliza append lo que hace es añadir sin reconstruir
	
	//Cap es la capacidad del slice
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	fmt.Println(slice)

}

