package main

import "fmt"

	type User struct {
		
		edad int
		nombre, apellido string
	
	}
	
	
	func (usuario User) nombre_completo() string {
		return usuario.nombre + " " + usuario.apellido
	}
	
	//IMPORTANTE: Esto no va a funcionar porque el user que entra en el método es una copia
	// NO se pueden crear metodos seters de esta forma
	/*func (usuario User) set_name(n string){
		usuario.nombre = n
	}*/
	
	// Para que funcione hay que pasarle el puntero ya que esto no es una copia si no que altera 
  // directamente el valor de lo que hay en el espacio de la memoria
  func (usuario *User) set_name(n string){
		usuario.nombre = n
	}

	func main(){

		user := new (User)
		user.nombre = ("Noel")
		user.apellido = ("Cerro")
		
		user.set_name("Otro Noel")
		
		fmt.Println(user.nombre_completo())
		
	}
