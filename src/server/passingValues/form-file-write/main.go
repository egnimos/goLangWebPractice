package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("server is started...")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		//fetch the file
		f, h, err := req.FormFile("f")
		errChecker(err, w)
		defer f.Close()

		fmt.Println("\nfile: ", f, "\nHeader: ", h, "\nerr", err)

		//read the file
		bs, err := ioutil.ReadAll(f)
		errChecker(err, w)

		s = string(bs)

		//create a file in the server...
		create, err := os.Create(h.Filename)
		errChecker(err, w)
		defer create.Close()

		_, err = create.Write(bs)
		errChecker(err, w)

	}

	tpl.ExecuteTemplate(w, "index.gohtml", s)
}

func errChecker(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}