package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

const SAMPLES = 200

func main() {
	var values [SAMPLES]float64

	for i := 0; i < SAMPLES; i++ {
		values[i] = float64(i / 2 % 16)
	}
	graph := asciigraph.Plot(values[:])

	fmt.Println(graph)
}
