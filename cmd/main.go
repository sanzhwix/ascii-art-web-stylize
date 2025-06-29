package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/print", printHandleFunc)

	http.ListenAndServe(":8080", nil)
}

func printHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}
