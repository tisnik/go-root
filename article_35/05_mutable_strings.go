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
// Demonstrační příklad číslo 5:
//    Alternativa k měnitelným řetězcům.

package main

import "fmt"

func main() {
	var s []byte = []byte("www.root.cz")

	fmt.Println(string(s))
	s[3] = '*'
	s[8] = '*'
	fmt.Println(string(s))
}
