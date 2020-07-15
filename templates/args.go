package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// func main() {
// 	name := os.Args[1]
// 	wd, _ := os.Getwd()

// 	tpl := fmt.Sprintf(`
// 		<!DOCTYPE html>
// 		<html lang="en">
// 			<head>
// 				<title>My website</title>
// 			</head>
// 			<body>
// 				<h1>My name is %s and this is my website</h1>
// 			</body>
// 		</html>
// 	`, name)

// 	nf, err := os.Create(wd + "/templates/static/index.html")
// 	if err != nil {
// 		log.Fatalln("Error creating index.html")
// 	}
// 	defer nf.Close()

// 	_, _ = io.Copy(nf, strings.NewReader(tpl))
// }
