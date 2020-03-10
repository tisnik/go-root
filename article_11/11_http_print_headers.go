// Seriál "Programovací jazyk Go"
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 11:
//    Vytištění hlavičky HTTP odpovědi

package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	response, err := http.Get("http://httpbin.org/uuid")
	if err != nil {
		println("Connection refused")
	}
	defer response.Body.Close()

	fmt.Printf("Status: %s\n", response.Status)
	fmt.Printf("Content length: %d\n", response.ContentLength)

	for name, headers := range response.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}
