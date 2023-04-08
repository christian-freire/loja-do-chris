package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{"Headset", "Qualidade alta", 250, 10},
		{"Mouse", "Ótimo DPI", 160, 15},
		{"Microfone", "Condensador", 300, 10},
		{"Teclado", "Mecânico", 180, 10},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
