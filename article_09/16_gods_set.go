// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Devátá část
//     Užitečné balíčky pro každodenní použití jazyka Go
//     https://www.root.cz/clanky/uzitecne-balicky-pro-kazdodenni-pouziti-jazyka-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_09/README.md
//
// Demonstrační příklad číslo 16:
//    Použití množin z knihovny GoDS.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/16_gods_set.html

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
