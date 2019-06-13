// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//
// Demonstrační příklad číslo 16:
//     	Opačně seřazení sekvence/řezu

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func printArray(prefix string, numbers []int) {
	var state string
	if sort.IntsAreSorted(numbers) {
		state = "sorted"
	} else {
		state = "unsorted"
	}
	fmt.Printf("%s variant of %s array: %v\n", prefix, state, numbers)
}

func main() {
	numbers := make([]int, 20)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Int() % 10
	}
	printArray("1st", numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	printArray("2nd", numbers)
}
