package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	wd, _ := os.Getwd()
	// Parse the template <tpl.gohtml> and put it into <tpl>
	tpl, err := template.ParseFiles(wd + "/templates/static/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Create a file called index.html
	nf, err := os.Create(wd + "/templates/static/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	// Execute the template inside the file
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
