package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, "Hello from TCP server \n")
		fmt.Fprintln(conn, "Everything's fine?")
		fmt.Fprintln(conn, "Well,I hope!")
		conn.Close()
	}

	// Testar usando "telnet localhost 8080"
}
