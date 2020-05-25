package main

import (
	"fmt"
	"net/http"
)

func handleRoute1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Route1")
}

func handleRoute2(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Route2")
}

func main() {

	muxServe := http.NewServeMux()
	muxServe.HandleFunc("/first", handleRoute1)
	muxServe.HandleFunc("/second/", handleRoute2)
	http.ListenAndServe(":8080", muxServe)

}
