package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	FirstName string
	LastName  string
	Items     []string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encoder", encoder)

	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// s := `<!DOCTYPE html>
	// 	<html lang="en">
	// 	<head>
	// 	<meta charset="UTF-8">
	// 	<title>FOO</title>
	// 	</head>
	// 	<body>
	// 	You are at foo
	// 	</body>
	// 	</html>`
	// w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "Niteesh",
		LastName:  "Dubey",
		Items:     []string{"devin", "kelvin kelin", "levis", "poppin", "jaeden", "gucci"},
	}

	sb, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(sb)
	// http.Redirect(w, req, "/", http.StatusSeeOther())
}

func encoder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		FirstName: "Jhon",
		LastName:  "Doe",
		Items:     []string{"bootup", "short", "leap", "cars", "home", "ladder"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Fatalln(err)
	}
}
