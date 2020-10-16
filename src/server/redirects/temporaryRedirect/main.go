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
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("your Request for foo function is: ", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("your request for bar method is: ", req.Method)

	//process of redirect
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req for barred function is: ", req.Method)

	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
