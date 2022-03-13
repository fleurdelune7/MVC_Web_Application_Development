package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNo = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./templates/home.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNo, nil)
}
