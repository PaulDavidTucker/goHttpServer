package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Starting the web server! Visit http://localhost:8080/quote to get a quote.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, quote.Go())
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Goodbye, World!")
	})

	http.ListenAndServe(":8080", nil)
}
