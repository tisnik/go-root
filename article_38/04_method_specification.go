// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá osmá část
//    Tvorba webových aplikací v Go s využitím projektu Gorilla web toolkit
//    https://www.root.cz/clanky/tvorba-webovych-aplikaci-v-go-s-vyuzitim-projektu-gorilla-web-toolkit/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté osmé části:
//    https://github.com/tisnik/go-root/blob/master/article_38/README.md

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// ADDRESS is a default address of HTTP server
const ADDRESS = ":8080"

var counter int
var mutex = &sync.Mutex{}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

func getCounterEndpoint(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(writer, "Counter: %d\n", counter)
	mutex.Unlock()
}

func setCounter(newValue int) {
	mutex.Lock()
	counter = newValue
	mutex.Unlock()
}

func setCounterEndpoint(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err == nil {
		number, err := strconv.ParseInt(string(body), 10, 0)
		if err == nil {
			setCounter(int(number))
			fmt.Fprintf(writer, "New counter value: %d\n", counter)
		} else {
			log.Printf("conversion failed for input string '%s'", string(body))
		}
	} else {
		log.Printf("request body is empty")
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", mainEndpoint).Methods("GET")
	router.HandleFunc("/counter", getCounterEndpoint).Methods("GET")
	router.HandleFunc("/counter", setCounterEndpoint).Methods("PUT")

	log.Println("Starting HTTP server at address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, router)
	if err != nil {
		log.Fatal("Unable to initialize HTTP server", err)
		os.Exit(2)
	}
}
