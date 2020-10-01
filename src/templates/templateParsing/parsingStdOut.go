package main

import (
	// "fmt"
	// "io"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	//Must function handle error on by own and return *template
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	//execute the template
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
