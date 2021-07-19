package main

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := []float64{3, 4, 5, 6, 9, 8, 5, math.Inf(1), 6, 10, 2, math.Inf(-1), 2, 5, 6}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}
