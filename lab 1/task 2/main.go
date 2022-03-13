package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/fp", func(w http.ResponseWriter, r *http.Request) {
		numOfBytes, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written %d", numOfBytes))
	})

	http.ListenAndServe(":8080", nil)

}
