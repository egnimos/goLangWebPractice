package main

import (
	"fmt"
	"net/http"
)

type home int

func (h home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>This is the main page by using the default handler</h1>")
}

func main() {
	var home home
	fmt.Println("Server is Started...")
	http.Handle("/", home)
	http.ListenAndServe(":8080", nil)
}
