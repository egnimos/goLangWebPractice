package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("index.gohtml").Funcs(fm).ParseFiles("index.gohtml"))
}

var fm = template.FuncMap{
	"add":    double,
	"square": square,
	"sqRoot": sqRoot,
}

func double(x float64) float64 {
	return x + x
}

func square(x float64) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	err := tpl.Execute(os.Stdout, 3.0)
	if err != nil {
		log.Fatalln(err)
	}
}
