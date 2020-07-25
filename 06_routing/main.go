package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Handle takes the pattern and something with the type Handler, it could
	// be a handler made with HandlerFunc, with takes a function with a specific
	// signature <func(ResponseWriter, *Request)>.
	http.Handle("/home", http.HandlerFunc(homeHandler))
	// HandleFunc take only a function with the same signature as above.
	// It CAN be an anonymous function or some with a name
	http.HandleFunc("/dog", dogHandler)

	_ = http.ListenAndServe(":8080", nil)
}

// Handler signature, called with the specific uri </home>.
func homeHandler(res http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(res, "This route is the /home route")
}

// Handler signature, called with the specific uri </dog>.
func dogHandler(writer http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(writer, "Hello dog")
}
