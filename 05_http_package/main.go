package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type handler struct{}

var onMemoryDataBase = map[string][]string{}
var tpl *template.Template

func init() {
	// Get the current directory
	wd, _ := os.Getwd()
	// Parse an specific template and store it on <tpl>
	tpl = template.Must(template.ParseFiles(wd + "/http_package/templates/index.gohtml"))
}

func main() {
	// Creates a generic handler, which have a method with the signature <func(ResponseWriter, *Request)>
	// so it inherit the <http.Handler> interface.
	h := handler{}

	// Listen all the methods in the handler <h>
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatalln(err)
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse the form, required to get the data in <request.Form>
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// Map each key and value inside <onMemoryDataBase>
	for k, v := range r.Form {
		onMemoryDataBase[k] = v
	}

	// Add to the response header <Some-Key>, which could be any key.
	w.Header().Add("Some-Key", "This is the value of a key")
	// Append to the response the StatusCode of 201, which is <created>
	w.WriteHeader(201)

	// Execute the template on ResponseWriter with <onMemoryDataBase> on it.
	err = tpl.ExecuteTemplate(w, "index.gohtml", onMemoryDataBase)
	if err != nil {
		log.Fatalln(err)
	}
}
