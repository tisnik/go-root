// Seriál "Programovací jazyk Go"
//
// Devátá část
//
// Demonstrační příklad číslo 1:
//    Import nainstalovaného balíčku.
//    Výpočet Levenštejnovy vzdálenosti dvou řetězců.

package main

import (
	"github.com/agext/levenshtein"
)

func main() {
	s1 := "Hello"
	s2 := "hello"
	println(levenshtein.Distance(s1, s2, nil))

	s1 = "Marta"
	s2 = "Markéta"
	println(levenshtein.Distance(s1, s2, nil))

	s1 = " foo"
	s2 = "jiný naprosto odlišný text nesouvisející s foo"
	println(levenshtein.Distance(s1, s2, nil))
}
