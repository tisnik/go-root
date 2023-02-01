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
		Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Partition(3).
		Map(func(value int) string {
			return strconv.Itoa(value)
		}).
		Reduce(func(x string, y string) string {
			return x + "," + y
		}).
		Map(func(x string) {
			println("joined:", x)
		}).
		Run()
}
