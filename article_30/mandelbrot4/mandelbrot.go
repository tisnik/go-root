// Seriál "Programovací jazyk Go"
//
// Třicátá část
//    Trasování a profilování aplikací naprogramovaných v Go (dokončení
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go-dokonceni/

package main

import (
	"context"
	"fmt"
	"log"
	"mandelbrot4/renderer"
	"os"
	"runtime/trace"
	"strconv"
)

func main() {
	f, err := os.Create("mandelbrot4.trace")
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

	if len(os.Args) < 4 {
		println("usage: ./mandelbrot width height maxiter")
		os.Exit(1)
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Improper width parameter: '%s'\n", os.Args[1])
		os.Exit(1)
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Improper height parameter: '%s'\n", os.Args[2])
		os.Exit(1)
	}

	maxiter, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Improper maxiter parameter: '%s'\n", os.Args[3])
		os.Exit(1)
	}

	ctx := context.Background()
	ctx, task := trace.NewTask(ctx, "renderFractal")
	defer task.End()
	renderer.Start(ctx, width, height, maxiter)
}
