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

func errorCheck(err error, text string) {
	if err != nil {
		log.Fatalln(text, " : ", err)
	}
}

func main() {
	sage := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sage)
	errorCheck(err, "error generated while executing the template")
}
