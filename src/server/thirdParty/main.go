package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl= template.Must(template.ParseGlob("template/*"))
}

func HandleError(w http.ResponseWriter, error error) {
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		log.Fatalln(error)
	}
}

func main() {
	mux := httprouter.New()
	mux.GET("/", Index)
	mux.GET("/about", About)
	mux.GET("/contact", Contact)
	mux.GET("/apply",  Apply)
	mux.POST("/apply", ApplyProcess)
	mux.GET("/user/:name", User)
	mux.GET("/blog/:category/:article", BlogRead)
	mux.POST("/blog/:category/:article", BlogWrite)

	//start the server
	fmt.Println("Server is Started....")
	http.ListenAndServe(":8080", mux)
}


