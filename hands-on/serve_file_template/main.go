package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog/", dog)

	_ = http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, _ *http.Request) {
	tpl, err := template.ParseFiles("templates/dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
