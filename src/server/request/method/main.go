package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type hotdog int

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func errorChecker(err error, text string) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	errorChecker(err, "parse form error")

	data := struct {
		Heading string
		Info    string
		Method  string
		Form    url.Values
	}{
		"Welcome to the form",
		"Here we test the request pointer and function and template",
		req.Method,
		req.Form,
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", data)
	errorChecker(err, "Error while execute")
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	fmt.Println("Server is Starting...")
	http.ListenAndServe(":8080", d)
}
