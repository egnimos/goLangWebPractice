package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"in": increment,
}

type hotdog int

func increment(i int) int {
	j := i
	k := j + 1
	return k
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	fmt.Println("Server is Started....")
	http.ListenAndServe(":8080", d)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	errorCheck(err)

	data := struct {
		Heading       string
		Info          string
		Method        string
		URL           *url.URL
		Form          url.Values
		InputName     string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		"Welcome To The Test",
		"This is the just simple test",
		req.Method,
		req.URL,
		req.Form,
		req.FormValue("fname"),
		req.Header,
		req.Host,
		req.ContentLength,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	errorCheck(err)
}
