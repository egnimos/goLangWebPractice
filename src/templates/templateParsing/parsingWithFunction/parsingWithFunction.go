package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	sage := []string{
		"Buddha",
		"Gandhi",
		"Martin luther king",
		"swami vivekanand",
		"Mahummad",
	}
	err := tpl.Execute(os.Stdout, sage)
	if err != nil {
		log.Fatalln(err)
	}
}
