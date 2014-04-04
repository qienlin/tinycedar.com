package handlers

import (
	"html/template"
	"log"
	"net/http"
	"templates"
)

type Handler struct {
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.String() {
	case "/":
		handleIndex(w, r)
	case "/login":
		handleLogin(w, r)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle index\n")
	var templ = template.Must(template.New("index").Parse(templates.IndexTemplate()))
	templ.Execute(w, r.FormValue("s"))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle Login\n")
	var templ = template.Must(template.New("home").Parse(templates.HomeTemplate()))
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "daniel" && username == password {
		templ.Execute(w, username)
	} else {
		handleIndex(w, r)
	}
}
