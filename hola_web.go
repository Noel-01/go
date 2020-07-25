package main

import(
	"net/http"
	"io"
	"fmt"
)


func main() {

	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Holaaaaaa")		
	})

	http.HandleFunc("/",handler)
	http.ListenAndServe(":8000", nil)

}

//Esta función siempre es así
func handler(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Hay una nueva petición")
	io.WriteString(w,"Hola mundo")
	

}
