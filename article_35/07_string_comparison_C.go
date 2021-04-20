// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá pátá část
//    Programovací jazyk Go pro skalní céčkaře (2.část)
//    https://www.root.cz/clanky/programovaci-jazyk-go-pro-skalni-ceckare-2-cast/
//
// Demonstrační příklad číslo 7C:
//     Porovnávání řetězců v Go s podporou češtiny.

package main

import "golang.org/x/text/collate"
import "golang.org/x/text/language"
import "fmt"

func main() {
	cl := collate.New(language.Czech)
	fmt.Println(cl.CompareString("aa", "ab"))
	fmt.Println(cl.CompareString("aa", "aa"))

	fmt.Println(cl.CompareString("e", "é"))
	fmt.Println(cl.CompareString("e", "ě"))
	fmt.Println(cl.CompareString("ě", "é"))

	fmt.Println(cl.CompareString("z", "é"))
	fmt.Println(cl.CompareString("z", "ě"))

	fmt.Println(cl.CompareString("h", "ch"))
	fmt.Println(cl.CompareString("ch", "i"))

	fmt.Println(cl.CompareString("Hrdina", "Chocholoušek"))
	fmt.Println(cl.CompareString("Crha", "Chocholoušek"))
}
