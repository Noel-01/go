package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {

	archivo, error := os.Open("./hola.txt")
	defer archivo.Close() //Importante cerrar el archivo, con "defer" lo hace automaticamente despues de un return o fin de run
	
	if error != nil {
		fmt.Println("Hubo un error")
	}
	
	scanner := bufio.NewScanner(archivo)
	
	var i int
	
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}
	
	

}
