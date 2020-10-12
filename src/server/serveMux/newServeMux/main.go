package main

import (
	"fmt"
	"net/http"
)

type route int

func (r route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Hello this page is running by the servemux<a href='/'>HOME</a></h1>")
}

type home int

func (h home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>navigate to the route page<a href='/route'>ROUTE</a></h1>")
}

func main() {
	var route route
	var home home
	serveMux := http.NewServeMux()
	serveMux.Handle("/", home)
	serveMux.Handle("/route", route)
	fmt.Println("Server is Started")
	http.ListenAndServe(":8080", serveMux)
}

