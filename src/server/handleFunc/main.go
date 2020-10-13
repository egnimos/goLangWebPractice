package main

import (
	"fmt"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "<h1>hello this is the first routing page<a href='/second'>SECOND</a></h1>")
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "<h1>hello this is the second rote page<a href='/'>FIRST</a></h1>")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/second", h2)
	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}
