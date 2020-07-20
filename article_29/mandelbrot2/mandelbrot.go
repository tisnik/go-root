// Seriál "Programovací jazyk Go"
//
// Dvacátá devátá část
//    Trasování a profilování aplikací naprogramovaných v Go
//    https://www.root.cz/clanky/trasovani-a-profilovani-aplikaci-naprogramovanych-v-go/

package main

import (
	"fmt"
	"log"
	"mandelbrot/renderer"
	"os"
	"runtime/pprof"
	"strconv"
)

func main() {
	f, err := os.Create("mandelbrot2.prof")
	if err != nil {
		log.Fatalf("failed to create profiler output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close profiler file: %v", err)
		}
	}()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalf("failed to start profle: %v", err)
	}
	defer pprof.StopCPUProfile()

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

	renderer.Start(width, height, maxiter)
}
