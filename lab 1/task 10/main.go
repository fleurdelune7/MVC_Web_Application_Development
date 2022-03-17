package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home.tpl")
}

func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "about.tpl")
}

func renderTemplate(w http.ResponseWriter, templateName string){
	parsedTemplate, errTmp := template.ParseFiles("./pages/"+templateName)
	if errTmp != nil {
		fmt.Println("Error template parsing", errTmp)
		return 
	}
	resolvedTemplate, errLay := parsedTemplate.ParseFiles("./layouts/base.layout.tpl")
	if errLay != nil {
		fmt.Println("Error layout parsing", errLay)
		return
	}

	errExe := resolvedTemplate.Execute(w, nil)
	if errExe != nil {
		fmt.Println("error executing template:", errExe)
	}
}


func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	
	http.ListenAndServe(portNumber, nil) 
}