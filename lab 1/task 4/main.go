package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello!")
	if err != nil {
		fmt.Println(err)
	}
}

func Warsaw(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "The Capital of Poland")
	if err != nil {
		fmt.Println(err)
	}
}

func Krakow(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Nice City")
	if err != nil {
		fmt.Println(err)
	}
}

func Gdansk(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Nearby the sea")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/warsaw", Warsaw)
	http.HandleFunc("/krakow", Krakow)
	http.HandleFunc("/gdanks", Gdansk)

	http.ListenAndServe(":8080", nil)
}
