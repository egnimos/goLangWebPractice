package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, "<img src='/toby.jpg'>")
}