package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

const SAMPLES = 64

func displayGraph(modulus int, slope int) {
	var values [SAMPLES]float64

	for i := 0; i < SAMPLES; i++ {
		values[i] = float64((i - SAMPLES/2) / slope % modulus)
	}
	graph := asciigraph.Plot(values[:])

	fmt.Println(graph)
}

func main() {
	const modulus = 16
	for slope := 1; slope <= 5; slope++ {
		displayGraph(modulus, slope)
		fmt.Println("\n")
	}
}
