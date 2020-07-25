package main

import "fmt"


	type Persona struct {

		nombre string
		
	}
	
	type Dummy struct {

		nombre string
		
	}


	type Tutor struct {

		Persona  //Campo anónimo
		Dummy    //Campo anónimo
		
	}



	func main(){

		tutor := Tutor{Persona{"Noel"}}
		
		// Se puede acceder al nombre de la structura interna sin tener que hacerla una referencia
		fmt.Println(tutor.nombre)
		
		// IMPORTANTE: puede provocar un error lo siguiente
		// Este error se debe a que las 2 structs tienen un campo nombre y no go no sabe cual hay que coger
    // tutor2 := Tutor{Persona{"Noel"}, Dummy{"Juan"}}
    // fmt.Println(tutor2.nombre)		


		// Solución
    // tutor2 := Tutor{Persona{"Noel"}, Dummy{"Juan"}}
    // fmt.Println(tutor2.Dummy.nombre)		



	}
