// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Desátá část
//     Užitečné balíčky pro každodenní použití Go (2), porovnání výkonnosti Go s céčkem
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-go-2-porovnani-vykonnosti-go-s-ceckem/
//
// Demonstrační příklad číslo 10:
//    Přístup k proměnným prostředí.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_10/10_env_var.html

package main

import (
	"fmt"
	"os"
)

func main() {
	envVars := os.Environ()

	for i, envVar := range envVars {
		fmt.Printf("%02d\t%s\n", i, envVar)
	}

	term, found := os.LookupEnv("TERM")
	if found {
		fmt.Printf("\n\n\nSelected TERM = %s", term)
	} else {
		fmt.Printf("\n\n\nThe TERM environment variable is not set")
	}
}
