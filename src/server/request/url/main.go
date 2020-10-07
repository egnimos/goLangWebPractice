package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

type hotdog int

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	errorChecker(err)

	data := struct {
		Method  string
		Heading string
		Info    string
		URL     *url.URL
		Form    url.Values
	}{
		req.Method,
		"Welcome to the Test From Golang",
		"Hello this is just a simple test",
		req.URL,
		req.Form,
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", data)
	errorChecker(err)
}

func main() {
	var d hotdog
	fmt.Println("Server is Starting....")
	err := http.ListenAndServe(":8080", d)
	errorChecker(err)
}

func errorChecker(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
