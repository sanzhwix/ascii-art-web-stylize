package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	print "ascii-art/art"
	validators "ascii-art/validation"
)

var Tpl *template.Template

// mux *http.ServeMux

type Data struct {
	Input  string
	Output string
	Full   bool
	Banner string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderErrorPage(w, "Error 404 Page not found", http.StatusNotFound)
		return
	}
	Tpl.ExecuteTemplate(w, "index.html", nil)
}

func PrintHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		font := r.FormValue("type")
		text := r.FormValue("text")
		maxInput := 100000

		if !validators.AsciiCharValidation(text) {
			RenderErrorPage(w, "Error 404: Only strandard ASCII charecters are allowed!", http.StatusBadRequest)
			return
		}
		if !validators.BannerValidity(font) {
			RenderErrorPage(w, "Error 500 Internal server error "+font+" banner was changed!", http.StatusInternalServerError)
			return
		}
		if len(text) > maxInput {
			RenderErrorPage(w, "Request is too large!", http.StatusBadRequest)
			return
		}
		if text == "" {
			RenderErrorPage(w, "Error 400 Bad request", http.StatusBadRequest)
			return
		}
		var buf bytes.Buffer
		print.Processing(text, font, &buf)
		res := Data{
			Input:  text,
			Output: buf.String(),
			Full:   true,
			Banner: font,
		}
		Tpl.ExecuteTemplate(w, "index.html", res)
	} else {
		RenderErrorPage(w, "Error 405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

type ErrorType struct {
	Msg string
}

func RenderErrorPage(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	data := ErrorType{Msg: msg}
	err := Tpl.ExecuteTemplate(w, "errorPage.html", data)
	if err != nil {
		fmt.Println("Error with executing error page html")
		return
	}
}
