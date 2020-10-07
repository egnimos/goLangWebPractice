package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseFiles("index.gohtml"))
// }

// func errorChecker(err error) {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Niteesh Dubey", "This is from Niteesh Dubey")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Server is Started and Header is Set</h1>")
}

func main() {
	var d hotdog
	fmt.Println("Server is started...")
	http.ListenAndServe(":8080", d)
}
