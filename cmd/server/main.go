package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		data := map[string][]Film{
			"Films": {
				{Title: "Liga da Justiça", Director: "Zack Snider"},
				{Title: "Era uma vez em Holliwood", Director: "Quentin Tarantino"},
				{Title: "Minha mãe é uma peça", Director: "Paulo Gustavo"},
			},
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		film := r.PostFormValue("film")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{
			Title:    film,
			Director: director,
		})
	})

	http.ListenAndServe(":8000", nil)
}
