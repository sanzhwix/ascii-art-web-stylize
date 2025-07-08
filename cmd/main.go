package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	print "ascii-art/art"
	validators "ascii-art/validation"
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
	if r.URL.Path != "/" {
		RenderErrorPage(w, "Error 404 Page not found", http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func printHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		font := r.FormValue("type")
		text := r.FormValue("text")

		if validators.AsciiCharValidation(text) {
			RenderErrorPage(w, "Error 404: Only strandard Ascii charecters are allowed!", http.StatusBadRequest)
			return
		}
		if validators.BannerValidity(font) {
			RenderErrorPage(w, "Error 500 Internal server error"+font+"was changed!", http.StatusInternalServerError)
			return
		}

		if font == "" || text == "" {
			RenderErrorPage(w, "Error 400 Bad request", http.StatusBadRequest)
			return
		}
		var buf bytes.Buffer
		print.Processing(text, font, &buf)
		res := Data{Output: buf.String(), Full: true}
		tpl.Execute(w, res)
	} else {
		RenderErrorPage(w, "Error 405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

type ErrorType struct {
	Msg string
}

func RenderErrorPage(w http.ResponseWriter, msg string, statusCode int) {
	t, err := template.ParseFiles("templates/errorPage.html")
	if err != nil {
		fmt.Println("Error parsing error page")
		return
	}
	w.WriteHeader(statusCode)
	data := ErrorType{Msg: msg}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error with executing error page html")
		return
	}
}

func main() {
	var err error

	tpl, err = template.ParseFiles("templates/index.html")
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
