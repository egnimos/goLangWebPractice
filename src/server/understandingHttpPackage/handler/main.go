package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello this is code is just for testing the handler and net http package")
}

func main() {
	var a hotdog
	fmt.Println("server is started....")
	http.ListenAndServe(":8080", a)
	defer fmt.Println("server is closed")
}
