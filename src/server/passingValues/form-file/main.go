package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
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
	var s string
	fmt.Println(req.Method)

	if req.Method == http.MethodPost {

		f, h, err := req.FormFile("f")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//print info
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr:", err)

		//read the file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalln(err)
		}
		 s = string(bs)
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	if err != nil {
		log.Fatalln(err)
	}
}
