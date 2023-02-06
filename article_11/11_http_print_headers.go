// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Jedenáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z jedenácté části:
//    https://github.com/tisnik/go-root/blob/master/article_11/README.md
//
// Demonstrační příklad číslo 11:
//    Vytištění hlavičky HTTP odpovědi
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_11/11_http_print_headers.html

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
