package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func errorChecker(err error, text string) {
	if err != nil {
		log.Fatalln(text, " : ", err)
	}
}

func main() {
	err := tpl.Execute(os.Stdout, 42)
	errorChecker(err, "error generated while execution")
}
