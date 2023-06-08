package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// datový typ, jehož prvky budou vypisovány v šabloně
type Role struct {
	Name       string
	Surname    string
	Popularity int
}

func sendStaticPage(writer http.ResponseWriter, filename string) {
	log.Printf("Sending static file %s", filename)

	body, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Fprint(writer, string(body))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
	}
}

func staticPage(filename string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		sendStaticPage(writer, filename)
	}
}

func rolesHandler(tmpl *template.Template, roles []Role) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		log.Printf("Application template for %d data records", len(roles))
		// aplikace šablony - přepis hodnot
		err := tmpl.Execute(writer, roles)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	}
}

func main() {
	const templateFilename = "html_template05.htm"

	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		{"Eliška", "Najbrtová", 4},
		{"Jenny", "Suk", 3},
		{"Anička", "Šafářová", 0},
		{"Sváťa", "Pulec", 3},
		{"Blažej", "Motyčka", 8},
		{"Eda", "Wasserfall", 0},
		{"Přemysl", "Hájek", 10},
	}

	log.Printf("Constructing template from file %s", templateFilename)

	// vytvoření nové šablony
	tmpl, err := template.ParseFiles(templateFilename)
	if err != nil {
		log.Fatalf("Template can't be constructed: %v", err)
		return
	}

	const address = ":8080"

	log.Printf("Starting server on address %s", address)

	http.HandleFunc("/", staticPage("index.html"))
	http.HandleFunc("/missing", staticPage("missing.html"))
	http.HandleFunc("/roles", rolesHandler(tmpl, roles))
	http.ListenAndServe(address, nil)
}
