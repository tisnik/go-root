// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_35/README.md
//
// Demonstrační příklad číslo 7B:
//     Porovnávání řetězců v Go bez podpory češtiny.

package main

import "fmt"
import "strings"

func main() {
	fmt.Println(strings.Compare("aa", "ab"))
	fmt.Println(strings.Compare("aa", "aa"))

	fmt.Println(strings.Compare("e", "é"))
	fmt.Println(strings.Compare("e", "ě"))
	fmt.Println(strings.Compare("ě", "é"))

	fmt.Println(strings.Compare("z", "é"))
	fmt.Println(strings.Compare("z", "ě"))

	fmt.Println(strings.Compare("h", "ch"))
	fmt.Println(strings.Compare("ch", "i"))

	fmt.Println(strings.Compare("Hrdina", "Chocholoušek"))
	fmt.Println(strings.Compare("Crha", "Chocholoušek"))
}
