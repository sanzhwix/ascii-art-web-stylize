package main

import (
	"ascii-art/internal/handlers"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	var err error
	handlers.Tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Something went wrong with templates parsing!")
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/print", handlers.PrintHandleFunc)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
