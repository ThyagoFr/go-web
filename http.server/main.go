package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Fprintln(w, "Responding the request...")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
