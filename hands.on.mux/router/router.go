package router

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

// DogRoute - Route
func DogRoute(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "dog route")
}

// CatRoute - Route
func CatRoute(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "cat route")
}

// MeRoute - Route
func MeRoute(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "thyago route")
}

// TemplateRoute - Route
func TemplateRoute(response http.ResponseWriter, request *http.Request) {
	path := "	"/home/thyagofr/Desktop/github/go.web/hands.on.mux/router/template.gohtml"
	tlp, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tlp.ExecuteTemplate(response, path, "thyago")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
