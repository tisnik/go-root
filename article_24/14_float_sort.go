// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//
// Demonstrační příklad číslo 14:
//     	Seřazení sekvence čísel s plovoucí řádovou čárkou

package main

import (
	"fmt"
	"math"
	"sort"
)

func printArray(prefix string, numbers []float64) {
	var state string
	if sort.Float64sAreSorted(numbers) {
		state = "sorted"
	} else {
		state = "unsorted"
	}
	fmt.Printf("%s variant of %s array: %v\n", prefix, state, numbers)
}

func main() {
	numbers := make([]float64, 20)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = math.Sin(float64(i) * math.Pi / float64(len(numbers)))
	}
	printArray("1st", numbers)

	sort.Float64s(numbers)
	printArray("2nd", numbers)
}
