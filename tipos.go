package main

import (
	"fmt"
	"strconv"
)

func main() {

	edad := "22"
/*	
	edad_str := strconv.Itoa(edad)
	
	fmt.Println("Mi edad es " + edad_str)
	*/
	
	edad_int,_ := strconv.Atoi(edad)
	
	fmt.Println(edad_int + 10)
	
}
	
	
