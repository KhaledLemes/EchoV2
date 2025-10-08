package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var templates *template.Template

func write(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}

func loadMainPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	echoed := r.FormValue("message")

	if echoed == "" {
		write(w, "You did not type anything D:")
		return
	}

	write(w, echoed)
}

func main() {
	router := mux.NewRouter()
	templates = template.Must(template.ParseGlob("*.html"))

	router.HandleFunc("/", loadMainPage).Methods(http.MethodGet)
	router.HandleFunc("/echo", echo).Methods(http.MethodPost)

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
