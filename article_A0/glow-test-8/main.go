package main

import (
	"flag"
	"strconv"

	_ "github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
)

func main() {
	flag.Parse()

	flow.
		New().
		TextFile("data.txt", 1).
		Map(func(line string) int {
			value, _ := strconv.Atoi(line)
			return value
		}).
		Reduce(func(x int, y int) int {
			return x + y
		}).
		Map(func(x int) {
			println("sum:", x)
		}).
		Run()
}
