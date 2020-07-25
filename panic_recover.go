package main

import(
	"fmt"
	"bufio"
	"os"
)

func readFile() bool {
	archivo, error := os.Open("./holaa.txt")
	
	defer func(){
		archivo.Close()
		fmt.Println("Defer")
		
		r := recover() // Detiene el panic
		fmt.Println(r)
	}()
	
	if error != nil {
		panic(error) //Lanza un error retornando todas las funciones y cierra la ejecución del script
	}
	
	scanner := bufio.NewScanner(archivo)
	
	var i int
	
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}
	
	fmt.Println("Nunca me ejecuto")
	
	return true
}

func executeReadFile() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func main() {
	executeReadFile()
	fmt.Println("Nunca me voy a imprimir")

	
}
