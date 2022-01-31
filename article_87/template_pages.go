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

func rolesHandler(templateFilename string, roles []Role) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Constructing template from file %s", templateFilename)

		// vytvoření nové šablony
		tmpl, err := template.ParseFiles(templateFilename)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Printf("Template can't be constructed: %v", err)
			return
		}

		log.Printf("Application template for %d data records", len(roles))
		// aplikace šablony - přepis hodnot
		err = tmpl.Execute(writer, roles)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
			return
		}
	}
}

func main() {
	// tyto hodnoty budou použity při aplikaci šablony
	roles := []Role{
		Role{"Eliška", "Najbrtová", 4},
		Role{"Jenny", "Suk", 3},
		Role{"Anička", "Šafářová", 0},
		Role{"Sváťa", "Pulec", 3},
		Role{"Blažej", "Motyčka", 8},
		Role{"Eda", "Wasserfall", 0},
		Role{"Přemysl", "Hájek", 10},
	}

	const address = ":8080"

	log.Printf("Starting server on address %s", address)

	http.HandleFunc("/", staticPage("index.html"))
	http.HandleFunc("/missing", staticPage("missing.html"))
	http.HandleFunc("/roles", rolesHandler("html_template05.htm", roles))
	http.ListenAndServe(address, nil)
}
