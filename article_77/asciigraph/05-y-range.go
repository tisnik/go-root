package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

const SAMPLES = 64

func displayGraph(modulus int) {
	var values [SAMPLES]float64

	for i := 0; i < SAMPLES; i++ {
		values[i] = float64(i % modulus)
	}
	graph := asciigraph.Plot(values[:])

	fmt.Println(graph)
}

func main() {
	for modulus := 2; modulus <= 64; modulus *= 2 {
		displayGraph(modulus)
		fmt.Println("\n")
	}
}
