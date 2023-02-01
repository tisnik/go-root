package main

import (
	"flag"

	"github.com/chrislusf/glow/flow"
)

func main() {
	flag.Parse()

	flow.
		New().
		Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Map(func(value int) int {
			return value * 2
		}).
		Reduce(func(x int, y int) int {
			return x + y
		}).
		Map(func(x int) {
			println("sum:", x)
		}).
		Run()
}
