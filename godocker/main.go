package main

import (
	"io"
	"log"
	"net/http"
)

// PingRoute - I dont know
func PingRoute(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Ping ...")
}

func main() {
	http.HandleFunc("/ping", PingRoute)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
