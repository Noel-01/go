package main

import "fmt"

type User interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct {
	nombre string
}

// Admin esta implementando la interfaz User con el metodo Permisos()
func (admin Admin) Permisos() int {
	return 5
}

// Admin esta implementando la interfaz User con el metodo Nombre()
func (admin Admin) Nombre() string {
	return admin.nombre
}

type Editor struct {
	nombre string
}

// Editor esta implementando la interfaz User con el metodo Permisos()
func (editor Editor) Permisos() int {
	return 3
}

// Editor esta implementando la interfaz User con el metodo Nombre()
func (editor Editor) Nombre() string {
	return editor.nombre
}

// User: se puede pasar cualquier struct que implemente los metodos de la interfaz User, 
//       en este caso lo hacen Admin y Editor
func auth(user User) string {
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	} else if permisos < 4 {
			return user.Nombre() + " tiene permisos de editor"
	}
	return ""
}

func main(){

	admin := Admin{"Noel"}
	editor := Editor{"Juan"}
	
	fmt.Println(auth(admin))	
	fmt.Println(auth(editor))	
	
	// También se puede hacer un array de strutcts que implementan la interfaz
	usuarios := []User{Admin{"Noel2"},Editor{"Juan2"}}
	
	//Recorrer el array con un bucle for, el '_' es el index pero aquí no me hace falta 
	for _,usuario := range usuarios {
		fmt.Println(auth(usuario))
	}
	

}
