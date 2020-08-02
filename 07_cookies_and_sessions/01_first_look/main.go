package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", foo)
	// Handles favicon
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	// Will throws an error if is not exists yet
	if err != nil {
		// Generate unique id with google/uuid package.
		id, _ := uuid.NewRandom()
		// Creates a pointer to a Cookie with type literals.
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			// Cookie not accessible with JavaScript, only with http.
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	// Print it back if it exists or newly created
	fmt.Println(cookie)
}
