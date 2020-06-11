// Seriál "Programovací jazyk Go"
//
// Dvacátá čtvrtá část
//     Kontejnery v základní knihovně programovacího jazyka Go
//     https://www.root.cz/clanky/kontejnery-v-zakladni-knihovne-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 13:
//     	Seřazení sekvence celých čísel

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

	sort.Ints(numbers)
	printArray("2nd", numbers)
}
