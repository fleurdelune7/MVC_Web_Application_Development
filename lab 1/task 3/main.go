package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		numOfBytes, err := fmt.Fprintf(w, "Hello!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written %d", numOfBytes))
	})

	http.HandleFunc("/warsaw", func(w http.ResponseWriter, r *http.Request) {
		numOfBytes, err := fmt.Fprintf(w, "The Capital of Poland")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written %d", numOfBytes))
	})

	http.HandleFunc("/krakow", func(w http.ResponseWriter, r *http.Request) {
		numOfBytes, err := fmt.Fprintf(w, "Nice City")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written %d", numOfBytes))
	})

	http.HandleFunc("/gdansk", func(w http.ResponseWriter, r *http.Request) {
		numOfBytes, err := fmt.Fprintf(w, "Nearby a sea")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written %d", numOfBytes))
	})

	http.ListenAndServe(":8080", nil)

}
