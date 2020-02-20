// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//
// Demonstrační příklad číslo 10:
//     Použití HTTP metody GET

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://httpbin.org/uuid")
	if err != nil {
		println("Connection refused")
	}
	defer response.Body.Close()

	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Content length: %d\n", response.ContentLength)

	rawBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println("Response body can't be read")
	} else {
		body := string(rawBody)
		println(body)
	}
}
