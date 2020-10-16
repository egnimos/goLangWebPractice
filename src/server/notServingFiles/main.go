package main

import (
	"fmt"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<img src='data:image/png;base64,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'>")
}

func main() {
	http.HandleFunc("/", dog)
	//start the server....
	fmt.Println("Server is Started....")
	http.ListenAndServe(":8080", nil)
}