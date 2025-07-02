package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func htmlHandle(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// func printHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {

// 		r.ParseForm()
// 		text := r.FormValue("text")

// 		tpl.ExecuteTemplate(w, "print.html", text)
// 	}
// }

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/", htmlHandle)

	// http.HandleFunc("/print", printHandleFunc)

	http.ListenAndServe(":8080", nil)
}
