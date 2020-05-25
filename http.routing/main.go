package main

import (
	"fmt"
	"net/http"
)

type server int

func (s server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/first":
		fmt.Fprintln(res, "Primeira rota...")
	case "/second":
		fmt.Fprintln(res, "Segunda rota...")
	}
}

func main() {
	var s server
	http.ListenAndServe(":8080", s)

}
