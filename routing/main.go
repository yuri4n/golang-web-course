package main

import (
	"fmt"
	"io"
	"net/http"
)

func homeHandler(res http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(res, "This route is the /home route")
}

func dogHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(writer, "Hello dog")
}

func main() {
	http.Handle("/home", http.HandlerFunc(homeHandler))
	http.HandleFunc("/dog", dogHandler)

	_ = http.ListenAndServe(":8080", nil)
}
