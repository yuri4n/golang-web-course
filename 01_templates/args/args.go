package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Take the second argument passed throw its execution.
	name := os.Args[1]
	wd, _ := os.Getwd()

	tpl := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<title>My website</title>
			</head>
			<body>
				<h1>My name is %s and this is my website</h1>
			</body>
		</html>
	`, name)

	// Create or replace a file called <index.html>.
	nf, err := os.Create(wd + "/templates/static/index.html")
	if err != nil {
		log.Fatalln("Error creating index.html")
	}
	defer nf.Close()

	// Copy the content of string <tpl> inside the file index.html.
	_, _ = io.Copy(nf, strings.NewReader(tpl))
}
