package main
//package handlerFunc
import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintln(w, "<h1>Hello this is the home page Navigate to the <a href='/about'>about</a> page or <a href='/contact'>contact</a> page")
}

func about(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Hey this is the egnimo is here and iam working on the golang and practicing how to create the server of your own if your navigate to other page there are different options below just click to go there happy coding<br><br><br><a href='/'>Home</a> page or <a href='/contact'>contact</a> page")
}

func contact(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Hello this is the home page Navigate to the <a href='/'>Home</a> page or <a href='/about'>about</a> page")
}




