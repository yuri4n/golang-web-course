package main

import "fmt"

func main() {
	name := "Julian Garzon"

	// Template made with string literal
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

	fmt.Println(tpl)
}
