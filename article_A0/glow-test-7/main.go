package main

import (
	"flag"
	"strconv"

	_ "github.com/chrislusf/glow/driver"
	"github.com/chrislusf/glow/flow"
)

func newRangeSlice(sliceLength int) []int {
	result := make([]int, sliceLength)
	for i := 0; i < sliceLength; i++ {
		result[i] = i + 1
	}
	return result
}

func compute(sliceLength int, partitions int) {
	flow.
		New().
		Slice(newRangeSlice(sliceLength)).
		Partition(partitions).
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

func main() {
	flag.Parse()

	compute(24, 1)
	compute(24, 2)
	compute(24, 3)
	compute(24, 4)
	compute(24, 6)
	compute(24, 8)
	compute(24, 12)
}
