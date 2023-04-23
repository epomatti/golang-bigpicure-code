package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Go web services")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":3000", nil)
}
