package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	print "ascii-art/art"
)

var (
	tpl *template.Template
	mux *http.ServeMux
)

type Data struct {
	Output string
	Full   bool
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func printHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		font := r.FormValue("type")
		text := r.FormValue("text")
		if font == "" || text == "" {
			w.WriteHeader(http.StatusBadRequest)
			RenderErrorPage()
			return
		}
		var buf bytes.Buffer
		print.Processing(text, font, &buf)
		res := Data{Output: buf.String(), Full: true}
		tpl.Execute(w, res)
	} else {
		RenderErrorPage()
	}
}

// function to use only one html error page
type ErrorType struct {
	Msg string
}

// If r.URL.PATH != "/" {
//}

func RenderErrorPage(w http.ResponseWriter, msg string, errCode int) {
	w.WriteHeader(errCode)
	err := ErrorType{Msg: msg}
	tpl.Execute(w, err)
}

func main() {
	var err error

	tpl, err = template.ParseFiles("templates/index.html", "errorPage.html")
	if err != nil {
		log.Fatal("Template parsing error:", err)
	}

	mux = http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/print", printHandleFunc)

	fmt.Println("Server is running here: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
