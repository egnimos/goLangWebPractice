package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<img src='/toby.jpg'>")
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "File not Found", 404)
		log.Fatalln(err)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
