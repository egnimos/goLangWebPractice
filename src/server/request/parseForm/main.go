package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var tpl *template.Template

type hotdog int

// type r *http.Request

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err, "error generates from form data")
	}

	fp := struct {
		Form    url.Values
		Heading string
		Info    string
	}{
		req.Form,
		"Welcome to the testing of the form data",
		"Hello in this we are passing the struct form of data function and passing request and response",
	}

	// fp := formPage{
	// 	Form:    req.Form,
	// 	Heading: "Welcome to the testing of the form data",
	// 	Info:    "Hello in this we are passing the struct form of data function and passing request and response",
	// }

	err = tpl.ExecuteTemplate(res, "index.gohtml", fp)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var d hotdog
	fmt.Println("Server is Starting....")
	http.ListenAndServe(":8080", d)
}
