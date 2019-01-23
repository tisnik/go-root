// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 16:
//    Použití množin z knihovny GoDS.

package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/linkedhashset"
)

const maxPrime = 1000

func printSet(set *linkedhashset.Set) {
	iterator := set.Iterator()
	for iterator.Next() {
		index, value := iterator.Index(), iterator.Value()
		fmt.Printf("%3d ", value)
		if index%10 == 9 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func main() {
	primes := linkedhashset.New()
	for n := 2; n < maxPrime; n++ {
		primes.Add(n)
	}

	for n := 2; n < maxPrime/2; n++ {
		if primes.Contains(n) {
			for k := 2 * n; k < maxPrime; k += n {
				primes.Remove(k)
			}
		}
	}
	printSet(primes)
}
