package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("../index.gohtml"))
}

func main() {
	nf, err := os.Create("index.gohtml")
	errChecker(err)

	err = tpl.Execute(nf, nil)
	errChecker(err)

}

func errChecker(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
