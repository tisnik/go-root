// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devadesátá druhá část
//    Knihovny s implementací generických datových typů pro programovací jazyk Go (2)
//    https://www.root.cz/clanky/knihovny-s-implementaci-generickych-datovych-typu-pro-programovaci-jazyk-go-2/
//
// Seznam příkladů z devadesáté druhé části:
//    https://github.com/tisnik/go-root/blob/master/article_92/README.md

package main

import (
	"fmt"

	"github.com/daichi-m/go18ds/sets/linkedhashset"
)

const maxPrime = 1000

func printSet(set *linkedhashset.Set[int]) {
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
	primes := linkedhashset.New[int]()
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
