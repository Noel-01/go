package main

import (
	"net/http"
	"encoding/json"
)

type Curso struct {
	
	/* IMPORTANTE: las variables en mayuscular son Publicas y en minusculas Privadas */
	// Con encoding/json puedo enviar a la web las variables con el nombre que quiera y en minuscula
	
	Title string `json:"titulo"`
	NumeroVideos int	`json:"numero_videos"`
	
}


type Cursos []Curso



func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	
		curso := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de Ruby", 20},
			Curso{"Curso de NodeJS", 40}, //Aquí el último también lleva coma ','
		}
		json.NewEncoder(w).Encode(curso)
	})
	
	http.ListenAndServe(":8000", nil)

}
