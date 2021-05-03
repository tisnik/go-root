// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá osmá část
//    Tvorba webových aplikací v Go s využitím projektu Gorilla web toolkit
//    https://www.root.cz/clanky/tvorba-webovych-aplikaci-v-go-s-vyuzitim-projektu-gorilla-web-toolkit/

package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

// ADDRESS is a default address of HTTP server
const ADDRESS = ":8080"

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func listAllPersonsEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "LIST ALL PERSONS\n")
}

func getPersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "GET PERSON\n")
}

func createPersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "CREATE PERSON\n")
}

func updatePersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "UPDATE PERSON\n")
}

func deletePersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "DELETE PERSON\n")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", mainEndpoint).Methods("GET")
	router.HandleFunc("/person", listAllPersonsEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/person/{id}", updatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/person/{id}", deletePersonEndpoint).Methods("DELETE")

	log.Println("Starting HTTP server at address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, router)
	if err != nil {
		log.Fatal("Unable to initialize HTTP server", err)
		os.Exit(2)
	}
}
