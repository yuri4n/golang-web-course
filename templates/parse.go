package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	wd, _ := os.Getwd()
	tpl, err := template.ParseFiles(wd + "/templates/static/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create(wd + "/templates/static/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
