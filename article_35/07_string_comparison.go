// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Seznam příkladů ze třicáté páté části:
//    https://github.com/tisnik/go-root/blob/master/article_35/README.md
//
// Demonstrační příklad číslo 7:
//     Porovnávání řetězců v Go.

package main

import "fmt"

func main() {
	fmt.Println("aa" < "ab")
	fmt.Println("aa" == "ab")

	fmt.Println("aa" < "aa")
	fmt.Println("aa" == "aa")

	fmt.Println("e" < "é")
	fmt.Println("e" < "ě")

	fmt.Println("z" < "é")
	fmt.Println("z" < "ě")

	fmt.Println("h" < "ch")
	fmt.Println("ch" < "i")

	fmt.Println("Hrdina" < "Chocholoušek")
	fmt.Println("Crha" < "Chocholoušek")
}
