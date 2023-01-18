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
// Demonstrační příklad číslo 1:
//    Import nainstalovaného balíčku.
//    Výpočet Levenštejnovy vzdálenosti dvou řetězců.
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_09/01_levenshtein.html

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
