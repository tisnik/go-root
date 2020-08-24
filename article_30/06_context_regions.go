// Seriál "Programovací jazyk Go"
//
// Třicátá část
//    Trasování a profilování aplikací naprogramovaných v Go (dokončení
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go-dokonceni/
//
// Demonstrační příklad číslo 6:
//    Transakce a regiony v trasovacím souboru.

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
	f, err := os.Create("trace3.trace")
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

	region1 := trace.StartRegion(ctx, "login")
	perform_login()
	region1.End()

	region2 := trace.StartRegion(ctx, "transactions")
	transaction("A")
	transaction("B")
	region2.End()

	region3 := trace.StartRegion(ctx, "logout")
	perform_logout()
	region3.End()

	task.End()
}
