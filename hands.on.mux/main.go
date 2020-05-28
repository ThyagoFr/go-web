package main

import (
	"log"
	"net/http"

	"github.com/thyagofr/go.web/handson/router"
)

func main() {

	http.HandleFunc("/dog", router.DogRoute)
	http.HandleFunc("/cat", router.CatRoute)
	http.HandleFunc("/me", router.MeRoute)
	http.HandleFunc("/template", router.TemplateRoute)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
