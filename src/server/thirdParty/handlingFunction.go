package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	fmt.Println("execution of the index template")
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func About(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	fmt.Println("execution of the about template")
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func Contact(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	fmt.Println("execution of the contact template")
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func Apply(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	fmt.Println("execution of the apply template")
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func ApplyProcess(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	fmt.Println("execution of the apply process template")
	err := tpl.ExecuteTemplate(w, "index.gohtml", "process is applied")
	HandleError(w, err)
}

func User(w http.ResponseWriter, req *http.Request, param httprouter.Params) {

}

func BlogRead(w http.ResponseWriter, req *http.Request, param httprouter.Params) {

}

func BlogWrite(w http.ResponseWriter, req *http.Request, param httprouter.Params) {

}

