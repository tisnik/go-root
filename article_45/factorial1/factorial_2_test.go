// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Čtyřicátá pátá část
//     Testování Go aplikací s využitím knihovny GΩmega a frameworku Ginkgo
//     https://www.root.cz/clanky/testovani-go-aplikaci-s-vyuzitim-knihovny-gomega-mega-a-frameworku-ginkgo/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze čtyřicáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_45/README.md

package main

import (
	"testing"
)

type factorialEntry struct {
	n        int64
	expected int64
}

func TestFactorial(t *testing.T) {
	var entries = []factorialEntry{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{9, 362880},
		{10, 3628800},
		{20, 2432902008176640000},
		{-1, 1},
	}
	for _, entry := range entries {
		computed := Factorial(entry.n)
		if computed != entry.expected {
			t.Errorf("%d! != %d, but %d", entry.n, computed, entry.expected)
		} else {
			t.Logf("factorial computer correctly for input %d", entry.n)
		}
	}
}
