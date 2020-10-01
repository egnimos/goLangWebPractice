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
	buddha := sage{
		"buddha",
		"The belief of no beliefs",
	}

	gandhi := sage{
		"gandhi",
		"Be the Change",
	}

	mlk := sage{
		"martin luther king",
		"Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		"jesus",
		"Love all",
	}

	muhammad := sage{
		"muhammad",
		"To Overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
