package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

const ADDRESS = ":8080"

var counter int
var mutex = &sync.Mutex{}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func counterEndpoint(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(writer, "Counter: %d\n", counter)
	mutex.Unlock()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", mainEndpoint)
	router.HandleFunc("/counter", counterEndpoint)

	log.Println("Starting HTTP server at address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, router)
	if err != nil {
		log.Fatal("Unable to initialize HTTP server", err)
		os.Exit(2)
	}
}
