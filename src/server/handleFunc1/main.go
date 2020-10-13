package main

import (
	"fmt"
	"net/http"
	
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	//start the server...
	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}