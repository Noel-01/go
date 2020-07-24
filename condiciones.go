package main

import "fmt"


func main() {
	
	x := 10
	y := 5
	
	/* No se puedo poner if () aquí no se ponen los parentesis*/
	if x > y {
		fmt.Printf("%d es mayor que %d \n", x,y)
	} else if x < y {
		fmt.Printf("%d es menor que %d \n", x,y)
	} else {
		fmt.Println("Son iguales")
	}
	
}
