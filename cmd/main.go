package main

import (
	print "ascii-art/art"
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
		print.Processing(text, w)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	tpl, _ = template.ParseGlob("templates/*.html")
	mux.HandleFunc("/", htmlHandle)

	mux.HandleFunc("/print", printHandleFunc)

	fmt.Println("Server is running here: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
