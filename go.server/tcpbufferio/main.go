package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handler(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	// Loop eterno pois o scanner ira ficar esperando dados ate que alguem feche a conexao...
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		// Printando o pacote http recebido...
		/*
			GET / HTTP/1.1
			Host: localhost:8080
			Connection: keep-alive
			Upgrade-Insecure-Requests: 1
			User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36 ...
		*/
	}
	defer conn.Close()
	fmt.Println("Fechando conexao...")
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn)
	}
}
