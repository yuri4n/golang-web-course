package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// Listen for any <tcp> connections on port <:8080>
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		// Handle each connection on a goroutine for concurrent requests.
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// Create an scanner based on the connection itself.
	scanner := bufio.NewScanner(conn)
	// Scan for each line.
	for i := 0; scanner.Scan(); i++ {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// Extract the uri field from the request header and store it on <uri> var
			uri := strings.Fields(ln)[1]
			// Call the mux method declared bellow and pass to it the connection itself and the
			// uri.
			mux(conn, uri)
		}
		// Breaks the loop if it found an empty line.
		if ln == "" {
			break
		}
	}

}

func mux(conn net.Conn, uri string) {
	var body string

	// Mux of request uri.
	switch uri {
	case "/":
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="eng"><head><title>My blog</title></head><body><h1>%s</h1></body></html>`, "Home")
	case "/posts":
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="eng"><head><title>My blog</title></head><body><h1>%s</h1></body></html>`, "Posts")
	default:
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="eng"><head><title>My blog</title></head><body><h1>%s</h1></body></html>`, "This route does not exists yet")
	}

	// NECESSARY HEADERS ON ANY HTTP RESPONSE
	// Write the header line into the connection, so now the client knows that it uses the
	// <HTTP> protocol and send to it the status code of <200, OK>.
	_, _ = fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	// Set the <Content-Length> header to the length of the <body>.
	_, _ = fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	// Tell to the connection that the content is <HTML>.
	_, _ = fmt.Fprint(conn, "Content-type: text/html \r\n")
	_, _ = fmt.Fprint(conn, "\r\n")
	// Write the body just right in the connection.
	_, _ = fmt.Fprint(conn, body)
}
