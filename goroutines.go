package main

import (
 "fmt"
 "time"
 "strings"
)

func mi_nombre_lentooo (nombre string) {

	letras := strings.Split(nombre, "")
	
	for _,letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	
}

func main() {

	// Simplemente con poner "go" delante de un metodo es concurrente (hilo)
	go mi_nombre_lentooo("Noel")
	fmt.Println("Que aburridooo")
	
	// Esto hace que no se acabe el programa hasta que no pulse una letra
	// para que le de tiempo al hilo a ejecutarse
	var wait string
	fmt.Scanln(&wait)

}
