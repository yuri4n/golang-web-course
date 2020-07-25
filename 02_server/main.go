package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen for a request
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	// Accept each connection in a infinite loop
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Open a goroutine to handle each connection
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		// Read each line send from the connection
		fmt.Println(ln)
		// Write "Hello world" in the client
		_, _ = fmt.Fprint(conn, "Hello world")
	}

	_ = conn.Close()
}
