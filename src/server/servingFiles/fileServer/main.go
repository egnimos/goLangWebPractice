package main

import (
	"fmt"
	//"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)

	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<img src='/toby.jpg'>")
}
