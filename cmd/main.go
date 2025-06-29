package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func htmlHandle(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func printHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		r.ParseForm()
		text := r.FormValue("text")

		fmt.Fprintf(w, "You submitted: %s", text)
	}
}

func main() {
	tpl, _ = template.ParseFiles("../index.html")
	http.HandleFunc("/", htmlHandle)

	http.HandleFunc("/print", printHandleFunc)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
