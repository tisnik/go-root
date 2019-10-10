// Seriál "Programovací jazyk Go"
//
// Třicátá část
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

	perform_login()
	transaction("A")
	transaction("B")
	perform_logout()

	task.End()
}
