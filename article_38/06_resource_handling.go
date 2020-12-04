// Seriál "Programovací jazyk Go"
//
// Třicátá osmá část
//    Tvorba webových aplikací v Go s využitím projektu Gorilla web toolkit
//    https://www.root.cz/clanky/tvorba-webovych-aplikaci-v-go-s-vyuzitim-projektu-gorilla-web-toolkit/

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

// ADDRESS is a default address of HTTP server
const ADDRESS = ":8080"

// Person structure represents information about person in IS
type Person struct {
	Firstname string `json:"firstname"`
	Surname   string `json:"lastname"`
}

var persons map[string]Person

func init() {
	persons = make(map[string]Person)
	persons["LT"] = Person{"Linus", "Torvalds"}
	persons["RP"] = Person{"Rob", "Pike"}
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func listAllPersonsEndpoint(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(persons)
}

func getPersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	person, found := persons[id]
	if found {
		json.NewEncoder(writer).Encode(person)
	} else {
		json.NewEncoder(writer).Encode(nil)
	}
}

func processPersonFromPayload(id string, request *http.Request) {
	var person Person
	err := json.NewDecoder(request.Body).Decode(&person)
	if err == nil {
		log.Println("JSON decoded")
		persons[id] = person
	} else {
		log.Println(err)
	}
}

func createPersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]

	_, found := persons[id]

	if !found {
		processPersonFromPayload(id, request)
	}
	json.NewEncoder(writer).Encode(persons)
}

func updatePersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "UPDATE PERSON\n")
	id := mux.Vars(request)["id"]

	_, found := persons[id]

	if found {
		processPersonFromPayload(id, request)
	}
	json.NewEncoder(writer).Encode(persons)
}

func deletePersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	_, found := persons[id]
	if found {
		delete(persons, id)
	}
	json.NewEncoder(writer).Encode(persons)
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
