package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	s := sage{
		"buddha",
		"the belif of no beliefs",
	}

	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalln(err, "execution error")
	}
}
