package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/dog":
		fmt.Fprintln(w, "<h1>Hello this is the DOGGY PAGE<a href='/cat'>CAT</a></h1>")
	case "/cat":
		fmt.Fprintln(w, "<h1>Hello this is the CATTY PAGE<a href='/dog'>DOG</a></h1>")
	case "/":
		fmt.Fprintln(w, "<ol><li><b><a href='/dog'>DOG</a></b></li><li><b><a href='/cat'>CAT</a></b></li></ol>")


	}
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}