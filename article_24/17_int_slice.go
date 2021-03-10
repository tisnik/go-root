// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá čtvrtá část
//     Kontejnery v základní knihovně programovacího jazyka Go
//     https://www.root.cz/clanky/kontejnery-v-zakladni-knihovne-programovaciho-jazyka-go/
//
// Demonstrační příklad číslo 17:
//     	Použití datového typu IntSlice

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func printIntSlice(numbers sort.IntSlice) {
	fmt.Printf("IntSlice: %v\n", numbers)
}

func main() {
	numbers := make([]int, 20)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Int() % 10
	}
	printIntSlice(numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	printIntSlice(numbers)
}
