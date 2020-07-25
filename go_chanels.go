package main

import (
 "fmt"
)

func main() {

  // Crear un canal de strings
	channel := make(chan string)
	
	// Función anónima
	go func (channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name //Se ENVIA "name" al canal "chanel"
		}
	}(channel)
	
	msg := <- channel //La variable msg RECIBE eventualmente información del channel
	
	fmt.Println("Estoy imprimiendo lo que recibí del canal: " + msg)
	
	msg = <- channel //La variable msg RECIBE eventualmente información del channel
	
	fmt.Println("Estoy imprimiendo lo que recibí del canal: " + msg)
	
}
