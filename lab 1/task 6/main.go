package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNo = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "about.html")
}

func renderTemplate(w http.ResponseWriter, r *http.Request, x string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + x)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNo, nil)
}
