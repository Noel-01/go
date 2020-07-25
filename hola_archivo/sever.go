package main

import(
	"net/http"
)

func main() {
	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.URL.Path)
		//http.ServeFile(w,r,"index.html")
		http.ServeFile(w,r,r.URL.Path[1:]) //Con esto puede acceder a todos lo archivos y los .html
	})
	
	http.ListenAndServe(":8000", nil)
	*/
	
	//Mejor hacerlo así y con las carpetas cradas
	fileServer := http.FileServer(http.Dir("public"))
	
	http.Handle("/", http.StripPrefix("/", fileServer))
	http.ListenAndServe(":8000", nil)
}

