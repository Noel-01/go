package main

import "fmt"

	type User struct {
		
		edad int
		nombre, apellido string
	
	}

func main(){

	var noel User
	fmt.Println(noel)
	
	fmt.Println(User{nombre: "Noel", apellido: "Cerro", edad: 30})
	
	usuario := User{nombre: "Noel", apellido: "Cerro", edad: 30}
	fmt.Println(usuario)
	fmt.Println(usuario.nombre)
	
	usuario.nombre = "Cambio nombre"
	fmt.Println(usuario.nombre)
	
	// El new no hace exactamente lo mismo que lo de arriba
	user := new (User)
	user.nombre = "Otro nombre"
	user.apellido = "Otro apellido"
	user.edad = 30
	
	fmt.Println(user)
	
}
