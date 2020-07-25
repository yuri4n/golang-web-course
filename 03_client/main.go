package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Dial to a tcp connection running on <localhost:8080>
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Write 'I dialed you' right on the connection
	_, _ = fmt.Fprintln(conn, "I dialed you.")
}
