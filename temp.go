package main

import (
	"fmt"
	"net/http"
	"temp/template"

	"github.com/a-h/templ"
)

const PORT = ":8080"

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("include_dir")))
	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	router.Handle("/tmpl", templ.Handler(template.Test()))

	fmt.Println(http.ListenAndServe(PORT, router))
}
