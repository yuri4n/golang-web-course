package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for i := 0; scanner.Scan(); i++ {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			uri := strings.Fields(ln)[1]

			mux(conn, uri)
		}
		if ln == "" {
			break
		}
	}

}

func mux(conn net.Conn, uri string) {
	var body string

	switch uri {
	case "/":
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="eng"><head><title>My blog</title></head><body><h1>%s</h1></body></html>`, "Home")
	case "/posts":
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="eng"><head><title>My blog</title></head><body><h1>%s</h1></body></html>`, "Posts")
	default:
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="eng"><head><title>My blog</title></head><body><h1>%s</h1></body></html>`, "This route does not exists yet")
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Lenght: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-type: text/html \r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
