package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art/internal/handlers"
)

func main() {
	var err error
	handlers.Tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Something went wrong with templates parsing!")
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "" {
			handlers.RenderErrorPage(w, "Fuzzing is restricted!", http.StatusNotFound)
			return
		}
		fs.ServeHTTP(w, r)
	})))

	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/ascii-art", handlers.PrintHandleFunc)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
