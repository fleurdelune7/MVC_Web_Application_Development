package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNo = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "home.tpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "about.tpl")

}

func renderTemplate(w http.ResponseWriter, r *http.Request, x string) {
	parsedTemplate, errTemp := template.ParseFiles("./pages/" + x)
	if errTemp != nil {
		fmt.Println("error parsing template:", errTemp)
	}
	resolvedTemplate, errLay := parsedTemplate.ParseFiles("./layouts/base.layout.tpl")
	if errLay != nil {
		fmt.Println("error parsing layout:", errLay)
	}
	errExe := resolvedTemplate.Execute(w, nil)
	if errExe != nil {
		fmt.Println("error executing template:", errExe)
	}

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNo, nil)
}
