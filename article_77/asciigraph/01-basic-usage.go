package main

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
)

func main() {
	data := []float64{3, 4, 5, 6, 9, 8, 5, 8, 6, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}
