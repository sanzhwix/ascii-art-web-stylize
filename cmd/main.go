package main

import (
	print "ascii-art/art"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type Data struct {
	Output string
	Full   bool
}

func htmlHandle(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func printHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		r.ParseForm()
		text := r.FormValue("text")
		var buf bytes.Buffer
		print.Processing(text, &buf)
		res := Data{Output: buf.String(), Full: true}
		tpl.Execute(w, res)
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
