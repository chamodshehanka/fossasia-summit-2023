package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, FOSSASIA 2023! \n")
	})

	fmt.Println("Starting FOSSASIA 2023 API Server...")
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}