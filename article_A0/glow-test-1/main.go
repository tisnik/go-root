package main

import (
	"flag"
	"fmt"

	"github.com/chrislusf/glow/flow"
)

func sum(x int, y int) int {
	return x + y
}

func printSum(x int) {
	fmt.Println("Sum:", x)
}

func main() {
	flag.Parse()

	flow.
		New().
		Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Reduce(sum).
		Map(printSum).Run()
}
