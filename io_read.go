package main

import(
	"io/ioutil"
	"fmt"
)

func main() {

	file_data,err := ioutil.ReadFile("./hola.txt") // "./" significa la misma carpeta desde donde ejecuto esto
	
	if err != nil {
		fmt.Println("Hubo un error")
	}
	
	fmt.Println(string(file_data))
	
	
}
