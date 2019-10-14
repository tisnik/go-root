package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const ADDRESS = ":8080"

type Person struct {
	Firstname string `json:"firstname"`
	Surname   string `json:"lastname"`
}

var persons map[int]Person

func init() {
	persons = make(map[int]Person)
	persons[0] = Person{"Linus", "Torvalds"}
	persons[1] = Person{"Rob", "Pike"}
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func listAllPersonsEndpoint(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(persons)
}

func retrieveIdRequestParameter(request *http.Request) int {
	id_var := mux.Vars(request)["id"]
	id, _ := strconv.ParseInt(id_var, 10, 0)
	return int(id)
}

func getPersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	id := retrieveIdRequestParameter(request)
	person, found := persons[id]
	if found {
		json.NewEncoder(writer).Encode(person)
	} else {
		json.NewEncoder(writer).Encode(nil)
	}
}

func processPersonFromPayload(id int, request *http.Request) {
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
	id := retrieveIdRequestParameter(request)

	_, found := persons[id]

	if !found {
		processPersonFromPayload(id, request)
	}
	json.NewEncoder(writer).Encode(persons)
}

func updatePersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "UPDATE PERSON\n")
	id := retrieveIdRequestParameter(request)

	_, found := persons[id]

	if found {
		processPersonFromPayload(id, request)
	}
	json.NewEncoder(writer).Encode(persons)
}

func deletePersonEndpoint(writer http.ResponseWriter, request *http.Request) {
	id := retrieveIdRequestParameter(request)

	_, found := persons[id]
	if found {
		delete(persons, id)
	}
	json.NewEncoder(writer).Encode(persons)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", mainEndpoint).Methods("GET")

	s := router.PathPrefix("/person").Subrouter()
	s.HandleFunc("", listAllPersonsEndpoint).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", getPersonEndpoint).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", createPersonEndpoint).Methods("POST").Headers("Content-Type", "application/json")
	s.HandleFunc("/{id:[0-9]+}", updatePersonEndpoint).Methods("PUT").Headers("Content-Type", "application/json")
	s.HandleFunc("/{id:[0-9]+}", deletePersonEndpoint).Methods("DELETE")

	log.Println("Starting HTTP server at address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, router)
	if err != nil {
		log.Fatal("Unable to initialize HTTP server", err)
		os.Exit(2)
	}
}
