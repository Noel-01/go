package main

import(
	"fmt"
	"bufio"
	"os"
)



func main() {

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Ingresa tu nombre: ")
	nombre, err := reader.ReadString('\n') //Lee hasta que le den al intro
	
	//nil es null
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}



/*
	var edad int
	var nombre string
	
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%v\n", &edad)
	fmt.Printf("Mi edad es %v\n", edad)
	
	fmt.Println("Coloca tu nombre: ")
	fmt.Scanf("%v\n", &nombre)
	fmt.Printf("Mi edad es %v\n", nombre)
*/


/*
		edad := 22
		nombre := "Noel"
		bander := true
		precio := 14.3
		
		fmt.Printf("Mi edad es %d y mi nombre es %v\n", edad, nombre)
		fmt.Printf("Mi nombre es %s\n", nombre)
		fmt.Printf("Este es verdadero %t\n", bander)
		fmt.Printf("El precio es %f\n", precio)
		*/
	}
	
