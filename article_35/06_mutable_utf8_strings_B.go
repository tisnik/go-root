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
// Demonstrační příklad číslo 6B:
//     Přístup k bajtům řetězce v UTF-8.

package main

import "fmt"

func main() {
	s := "[шщэюя]"

	var r []rune = []rune(s)

	fmt.Println(string(r))

	r[0] = '*'

	r[3] = '-'

	r[6] = '*'
	fmt.Println(string(r))
}
