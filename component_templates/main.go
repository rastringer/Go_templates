package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseGlob("templates/*.html"))

func displayHome(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "home", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func displayAbout(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "about", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func displayWorks(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "works", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func displayContact(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "contact", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", displayHome)
	http.HandleFunc("/about", displayAbout)
	http.HandleFunc("/works", displayWorks)
	http.HandleFunc("/contact", displayContact)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.ListenAndServe(":8080", nil)
}
