package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type handler struct{}

var onMemoryDataBase = map[string][]string{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	for k, v := range r.Form {
		onMemoryDataBase[k] = v
	}

	w.Header().Add("Some-Key", "This is the value of a key")
	w.WriteHeader(201)

	err = tpl.ExecuteTemplate(w, "index.gohtml", onMemoryDataBase)
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	wd, _ := os.Getwd()
	tpl = template.Must(template.ParseFiles(wd + "/http_package/templates/index.gohtml"))
}

func main() {
	h := handler{}

	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatalln(err)
	}
}
