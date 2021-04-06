// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá část
//    Trasování a profilování aplikací naprogramovaných v Go (dokončení
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go-dokonceni/
//
// Demonstrační příklad číslo 5:
//    Zpracování kontextu v trasovacích informacích.

package main

import (
	"context"
	"log"
	"os"
	"runtime/trace"
)

func performLogin() {
	log.Println("login")
}

func performLogout() {
	log.Println("logout")
}

func transaction(typ string) {
	log.Printf("transaction '%s'\n", typ)
}

func main() {
	f, err := os.Create("trace2.trace")
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

	ctx := context.Background()
	ctx, task := trace.NewTask(ctx, "transactionTask")

	performLogin()
	transaction("A")
	transaction("B")
	performLogout()

	task.End()
}
