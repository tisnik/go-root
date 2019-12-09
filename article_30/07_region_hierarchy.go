// Seriál "Programovací jazyk Go"
//
// Třicátá část
//
// Demonstrační příklad číslo 7:
//    Hierarchie regionů v souboru s trasovacími informacemi.

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

func transaction(ctx context.Context, typ string) {
	region := trace.StartRegion(ctx, "transaction")
	trace.Logf(ctx, "transaction", "type %s", typ)
	log.Printf("transaction '%s'\n", typ)
	region.End()
}

func main() {
	f, err := os.Create("trace4.trace")
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
	performLogin()
	region1.End()

	region2 := trace.StartRegion(ctx, "transactions")
	transaction(ctx, "A")
	transaction(ctx, "B")
	region2.End()

	region3 := trace.StartRegion(ctx, "logout")
	performLogout()
	region3.End()

	task.End()
}
