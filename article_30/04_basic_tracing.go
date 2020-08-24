// Seriál "Programovací jazyk Go"
//
// Třicátá část
//    Trasování a profilování aplikací naprogramovaných v Go (dokončení
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go-dokonceni/
//
// Demonstrační příklad číslo 4:
//     Vytvoření souboru s trasovacími informacemi.

package main

import (
	"log"
	"os"
	"runtime/trace"
)

func perform_login() {
	log.Println("login")
}

func perform_logout() {
	log.Println("logout")
}

func transaction(typ string) {
	log.Printf("transaction '%s'\n", typ)
}

func main() {
	f, err := os.Create("trace1.trace")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	perform_login()
	transaction("A")
	transaction("B")
	perform_logout()
}
