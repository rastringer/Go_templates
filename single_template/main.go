package main

import (
	"html/template"
	"net/http"
)

// Page is the page
type Page struct {
	Name  string
	Quote string
}

var t = template.Must(template.ParseFiles("templates/base.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := t.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func displayHome(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Name:  "The finest swordsmen and friends.",
		Quote: "All for one and one for all!",
	}
	renderTemplate(w, "base", p)
}

func displayDArtagnan(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Name:  "D'Artagnan",
		Quote: "Your Grace’s pardon, I had forgotten: England is an island. How do I get off?",
	}
	renderTemplate(w, "base", p)
}

func displayAthos(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Name:  "Athos",
		Quote: "You will find, young man, that the future looks rosiest through the bottom of a glass.",
	}
	renderTemplate(w, "base", p)
}

func displayPorthos(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Name:  "Porthos",
		Quote: "I fight just as well with my left hand. If you find that this places you at a disadvantage, I do apologize.",
	}
	renderTemplate(w, "base", p)
}

func displayAramis(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Name:  "Aramis",
		Quote: "Let us pray for our sins...on second thought, God's often busy.",
	}
	renderTemplate(w, "base", p)
}

func main() {
	http.HandleFunc("/", displayHome)
	http.HandleFunc("/dartagnan", displayDArtagnan)
	http.HandleFunc("/athos", displayAthos)
	http.HandleFunc("/porthos", displayPorthos)
	http.HandleFunc("/aramis", displayAramis)
	http.ListenAndServe(":3000", nil)
}
