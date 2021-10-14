// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvanáctá část
//     Vývoj síťových aplikací v programovacím jazyku Go (pokračování)
//     https://www.root.cz/clanky/vyvoj-sitovych-aplikaci-v-programovacim-jazyku-go-pokracovani/
//
// Demonstrační příklad číslo 9:
//    Druhá varianta webové aplikace pro výpočet faktoriálu
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_12/09_factorial_compute.html

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// FactorialPageDynContent holds all required information for factorial page template
type FactorialPageDynContent struct {
	N      int64
	Result int64
}

// Factorial computes factorial for a given n that must be non negative integer
func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	n, err := strconv.ParseInt(request.FormValue("n"), 10, 64)
	if err != nil {
		n = 0
	}

	t, err := template.ParseFiles("factorial_compute.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "Not found!")
		return
	}

	dynData := FactorialPageDynContent{N: n, Result: Factorial(n)}
	err = t.Execute(writer, dynData)
	if err != nil {
		println("Error executing template")
	}
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
