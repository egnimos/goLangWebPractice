package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	err := tpl.ExecuteTemplate(w, "index.gohtml", v)
	if err != nil {
		log.Fatalln(err)
	}
}
