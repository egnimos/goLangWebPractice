package main

import (
	"fmt"
	"net/http"
	//"io"
	//"os"
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
	//f, err := os.Open("../toby.jpg")
	//if err != nil {
	//	http.Error(w, err.Error(), 404)
	//	return
	//}
	http.ServeFile(w, req, "../toby.jpg")
}
